// Package github provides an OIDC token source backed by the GitHub Actions ID
// token endpoint.
package github

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/tidwall/gjson"

	"github.com/formalco/go-sdk/v3/oidc"
)

const (
	requestURLEnv        = "ACTIONS_ID_TOKEN_REQUEST_URL"
	requestTokenEnv      = "ACTIONS_ID_TOKEN_REQUEST_TOKEN"
	refreshSkew          = 30 * time.Second
	maxResponseBodyBytes = 1 << 20
)

type tokenSource struct {
	requestURL   string
	requestToken string
	client       *http.Client

	mu     sync.Mutex
	cached oidc.Token
}

// NewTokenSource returns a TokenSource that requests GitHub Actions ID tokens
// for audience from ACTIONS_ID_TOKEN_REQUEST_URL using httpClient. GitHub only
// sets the request variables for jobs granted the id-token: write permission.
// The request token stays valid for the whole job, so new ID tokens can be
// fetched as old ones expire.
func NewTokenSource(audience string, httpClient *http.Client) (oidc.TokenSource, error) {
	rawRequestURL := strings.TrimSpace(os.Getenv(requestURLEnv))
	requestToken := strings.TrimSpace(os.Getenv(requestTokenEnv))
	if rawRequestURL == "" || requestToken == "" {
		return nil, fmt.Errorf(
			"oidc: %s and %s are required; grant the GitHub Actions job the id-token: write permission",
			requestURLEnv,
			requestTokenEnv,
		)
	}
	if err := oidc.ValidateAudience(audience); err != nil {
		return nil, err
	}
	requestURL, err := url.Parse(rawRequestURL)
	if err != nil {
		return nil, fmt.Errorf("oidc: parse %s: %w", requestURLEnv, err)
	}
	query := requestURL.Query()
	query.Set("audience", audience)
	requestURL.RawQuery = query.Encode()

	return &tokenSource{
		requestURL:   requestURL.String(),
		requestToken: requestToken,
		client:       httpClient,
	}, nil
}

func (s *tokenSource) Token(ctx context.Context) (token oidc.Token, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.cached.JWT != "" && time.Now().Add(refreshSkew).Before(s.cached.Expiry) {
		return s.cached, nil
	}
	defer func() {
		if err != nil && s.cached.JWT != "" && time.Now().Before(s.cached.Expiry) {
			token, err = s.cached, nil
		}
	}()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, s.requestURL, nil)
	if err != nil {
		return oidc.Token{}, fmt.Errorf("create GitHub Actions OIDC request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+s.requestToken)
	req.Header.Set("Accept", "application/json")

	response, err := s.client.Do(req)
	if err != nil {
		return oidc.Token{}, fmt.Errorf("fetch GitHub Actions OIDC token: %w", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(io.LimitReader(response.Body, maxResponseBodyBytes))
	if err != nil {
		return oidc.Token{}, fmt.Errorf("read GitHub Actions OIDC response: %w", err)
	}
	if response.StatusCode != http.StatusOK {
		return oidc.Token{}, fmt.Errorf(
			"fetch GitHub Actions OIDC token: HTTP %d: %s",
			response.StatusCode,
			strings.TrimSpace(string(body)),
		)
	}

	jwt := strings.TrimSpace(gjson.GetBytes(body, "value").String())
	if jwt == "" {
		return oidc.Token{}, errors.New("GitHub Actions OIDC response returned an empty token")
	}
	parts := strings.Split(jwt, ".")
	if len(parts) != 3 {
		return oidc.Token{}, errors.New("GitHub Actions OIDC response returned a malformed JWT")
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return oidc.Token{}, errors.New("GitHub Actions OIDC response returned a malformed JWT")
	}
	expiresAt := gjson.GetBytes(payload, "exp").Int()
	if expiresAt == 0 {
		return oidc.Token{}, errors.New("GitHub Actions OIDC response returned a JWT without a valid expiry")
	}
	expiry := time.Unix(expiresAt, 0).UTC()
	if !time.Now().Before(expiry) {
		return oidc.Token{}, errors.New("GitHub Actions OIDC response returned an expired token")
	}
	s.cached = oidc.Token{JWT: jwt, Expiry: expiry}
	return s.cached, nil
}
