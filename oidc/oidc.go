package oidc

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/samber/mo"

	"github.com/formalco/typeid"
)

const integrationOIDCType = "integrationoidc"

// AudiencePrefix is the required prefix for Formal OIDC integration audiences.
const AudiencePrefix = "oidc.formal.ai/"

// IntegrationIDHeader selects an OIDC integration for providers with a fixed audience.
const IntegrationIDHeader = "X-Formal-OIDC-Integration-Id"

// Token contains an OIDC JWT and its request authentication metadata.
type Token struct {
	JWT                 string
	Expiry              time.Time
	HeaderIntegrationID mo.Option[string]
}

// TokenSource returns OIDC tokens accepted by Formal.
type TokenSource interface {
	Token(context.Context) (Token, error)
}

type staticTokenSource struct {
	jwt string
}

// Static returns a TokenSource that always returns jwt with an unknown expiry.
func Static(jwt string) TokenSource {
	return staticTokenSource{jwt: jwt}
}

func (s staticTokenSource) Token(context.Context) (Token, error) {
	if strings.TrimSpace(s.jwt) == "" {
		return Token{}, errors.New("oidc: token must not be empty")
	}
	return Token{JWT: s.jwt}, nil
}

type fileTokenSource struct {
	path string
}

// File returns a TokenSource that reads a JWT from path on every call, so
// tokens rotated in place (such as projected Kubernetes ServiceAccount tokens)
// are picked up without a restart.
func File(path string) TokenSource {
	return fileTokenSource{path: path}
}

func (s fileTokenSource) Token(context.Context) (Token, error) {
	contents, err := os.ReadFile(s.path)
	if err != nil {
		return Token{}, fmt.Errorf("oidc: read token file: %w", err)
	}
	jwt := strings.TrimSpace(string(contents))
	if jwt == "" {
		return Token{}, errors.New("oidc: token file must not be empty")
	}
	return Token{JWT: jwt}, nil
}

// ValidateAudience validates a Formal OIDC integration audience.
func ValidateAudience(audience string) error {
	if audience == "" {
		return errors.New("oidc: audience must not be empty")
	}
	if !strings.HasPrefix(audience, AudiencePrefix) {
		return errors.New("oidc: audience must be a Formal OIDC integration audience")
	}

	integrationID := strings.TrimPrefix(audience, AudiencePrefix)
	parsed, err := typeid.FromString(integrationID)
	if err != nil || parsed.Type() != integrationOIDCType {
		return errors.New("oidc: audience must contain a valid OIDC integration ID")
	}
	return nil
}
