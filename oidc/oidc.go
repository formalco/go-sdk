package oidc

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/formalco/typeid"
)

const integrationOIDCType = "integrationoidc"

// AudiencePrefix is the required prefix for Formal OIDC integration audiences.
const AudiencePrefix = "oidc.formal.ai/"

// Token is an OIDC JWT and its expiration time.
type Token struct {
	JWT    string
	Expiry time.Time
}

// TokenSource returns Formal-audience OIDC tokens.
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
