package cursor

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/formalco/go-sdk/v3/oidc"
)

const (
	agentSocketEnv       = "CURSOR_AGENT_SOCKET"
	tokenURL             = "http://cursor-agent/v1/tokens/oidc"
	refreshSkew          = 30 * time.Second
	maxResponseBodyBytes = 4 << 20
)

type tokenSource struct {
	audience string
	client   *http.Client

	mu     sync.Mutex
	cached oidc.Token
}

type mintRequest struct {
	Audience string `json:"aud"`
}

type mintResponse struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

// NewTokenSource returns a TokenSource backed by the Cursor Cloud Agent identity
// socket in CURSOR_AGENT_SOCKET.
func NewTokenSource(audience string) (oidc.TokenSource, error) {
	socketPath := strings.TrimSpace(os.Getenv(agentSocketEnv))
	if socketPath == "" {
		return nil, errors.New("oidc: Cursor agent socket path is required")
	}
	if err := oidc.ValidateAudience(audience); err != nil {
		return nil, err
	}

	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.Proxy = nil
	transport.DialContext = func(ctx context.Context, _, _ string) (net.Conn, error) {
		return (&net.Dialer{}).DialContext(ctx, "unix", socketPath)
	}
	return &tokenSource{
		audience: audience,
		client:   &http.Client{Transport: transport},
	}, nil
}

func (s *tokenSource) Token(ctx context.Context) (oidc.Token, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.cached.JWT != "" && time.Now().Add(refreshSkew).Before(s.cached.Expiry) {
		return s.cached, nil
	}

	token, err := s.mint(ctx)
	if err != nil {
		if s.cached.JWT != "" && time.Now().Before(s.cached.Expiry) {
			return s.cached, nil
		}
		return oidc.Token{}, err
	}
	s.cached = token
	return token, nil
}

func (s *tokenSource) mint(ctx context.Context) (oidc.Token, error) {
	body, err := json.Marshal(mintRequest{Audience: s.audience})
	if err != nil {
		return oidc.Token{}, fmt.Errorf("encode Cursor OIDC request: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, tokenURL, bytes.NewReader(body))
	if err != nil {
		return oidc.Token{}, fmt.Errorf("create Cursor OIDC request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	response, err := s.client.Do(req)
	if err != nil {
		return oidc.Token{}, fmt.Errorf("mint Cursor OIDC token: %w", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		message, _ := io.ReadAll(io.LimitReader(response.Body, maxResponseBodyBytes))
		return oidc.Token{}, fmt.Errorf(
			"mint Cursor OIDC token: HTTP %d: %s",
			response.StatusCode,
			strings.TrimSpace(string(message)),
		)
	}

	var minted mintResponse
	if err := json.NewDecoder(io.LimitReader(response.Body, maxResponseBodyBytes)).Decode(&minted); err != nil {
		return oidc.Token{}, fmt.Errorf("decode Cursor OIDC response: %w", err)
	}
	if strings.TrimSpace(minted.Token) == "" {
		return oidc.Token{}, errors.New("Cursor OIDC response returned an empty token")
	}
	expiry := time.Unix(minted.ExpiresAt, 0).UTC()
	if !time.Now().Before(expiry) {
		return oidc.Token{}, errors.New("Cursor OIDC response returned an expired token")
	}
	return oidc.Token{JWT: minted.Token, Expiry: expiry}, nil
}
