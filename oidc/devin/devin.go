package devin

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/formalco/go-sdk/v3/oidc"
)

const (
	defaultExecutable = "devin-oidc"
	refreshSkew       = 30 * time.Second
)

type tokenSource struct {
	audience   string
	executable string

	mu     sync.Mutex
	cached oidc.Token
}

type jwtClaims struct {
	ExpiresAt int64 `json:"exp"`
}

// NewTokenSource returns a TokenSource backed by devin-oidc from PATH.
func NewTokenSource(audience string) (oidc.TokenSource, error) {
	if err := oidc.ValidateAudience(audience); err != nil {
		return nil, err
	}
	executable, err := exec.LookPath(defaultExecutable)
	if err != nil {
		return nil, fmt.Errorf("find %s: %w", defaultExecutable, err)
	}
	return &tokenSource{audience: audience, executable: executable}, nil
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
	command := exec.CommandContext(ctx, s.executable, "token", "--audience", s.audience)
	output, err := command.Output()
	if err != nil {
		if exitErr, ok := errors.AsType[*exec.ExitError](err); ok && len(exitErr.Stderr) > 0 {
			return oidc.Token{}, fmt.Errorf(
				"mint Devin OIDC token: %w: %s",
				err,
				strings.TrimSpace(string(exitErr.Stderr)),
			)
		}
		return oidc.Token{}, fmt.Errorf("mint Devin OIDC token: %w", err)
	}
	jwt := strings.TrimSpace(string(output))
	if jwt == "" {
		return oidc.Token{}, errors.New("Devin OIDC command returned an empty token")
	}
	expiry, err := expiryFromJWT(jwt)
	if err != nil {
		return oidc.Token{}, err
	}
	if !time.Now().Before(expiry) {
		return oidc.Token{}, errors.New("Devin OIDC command returned an expired token")
	}
	return oidc.Token{JWT: jwt, Expiry: expiry}, nil
}

func expiryFromJWT(jwt string) (time.Time, error) {
	parts := strings.Split(jwt, ".")
	if len(parts) != 3 {
		return time.Time{}, errors.New("Devin OIDC command returned a malformed JWT")
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return time.Time{}, errors.New("Devin OIDC command returned a malformed JWT")
	}
	var claims jwtClaims
	if err := json.Unmarshal(payload, &claims); err != nil || claims.ExpiresAt == 0 {
		return time.Time{}, errors.New("Devin OIDC command returned a JWT without a valid expiry")
	}
	return time.Unix(claims.ExpiresAt, 0).UTC(), nil
}
