package aws

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	awssdk "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/smithy-go"

	"github.com/formalco/go-sdk/v3/oidc"
)

const (
	signingAlgorithm        = "ES384"
	tokenDurationSeconds    = int32(300)
	minTokenDurationSeconds = int32(60)
	sessionExpirySafetySkew = 2 * time.Second
	tokenRenewalPercent     = 60
)

type tokenSource struct {
	audience    string
	client      *sts.Client
	credentials awssdk.CredentialsProvider

	mu       sync.Mutex
	cached   oidc.Token
	cachedAt time.Time
}

// NewTokenSource returns a Formal-audience TokenSource backed by AWS STS GetWebIdentityToken.
// The caller owns AWS config, region, and credentials on client.
func NewTokenSource(client *sts.Client, audience string) (oidc.TokenSource, error) {
	if client == nil {
		return nil, errors.New("oidc: STS client is required")
	}
	if err := oidc.ValidateAudience(audience); err != nil {
		return nil, err
	}
	credentials := client.Options().Credentials
	if credentials == nil {
		return nil, errors.New("oidc: STS client is missing credentials")
	}

	return &tokenSource{
		audience:    audience,
		client:      client,
		credentials: credentials,
	}, nil
}

func (s *tokenSource) Token(ctx context.Context) (oidc.Token, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.cached.JWT != "" && !s.cached.Expiry.IsZero() && !s.cachedAt.IsZero() {
		lifetime := s.cached.Expiry.Sub(s.cachedAt)
		renewAt := s.cachedAt.Add(lifetime * tokenRenewalPercent / 100)
		if lifetime > 0 && time.Now().UTC().Before(renewAt) {
			return s.cached, nil
		}
	}

	token, err := s.mint(ctx)
	if err != nil {
		if s.cached.JWT != "" && !s.cached.Expiry.IsZero() && time.Now().UTC().Before(s.cached.Expiry) {
			return s.cached, nil
		}
		return oidc.Token{}, err
	}
	s.cached = token
	s.cachedAt = time.Now().UTC()
	return token, nil
}

func (s *tokenSource) mint(ctx context.Context) (oidc.Token, error) {
	creds, err := s.credentials.Retrieve(ctx)
	if err != nil {
		return oidc.Token{}, fmt.Errorf("retrieve AWS credentials: %w", err)
	}
	duration := durationSecondsForSession(creds, time.Now())
	if duration == 0 {
		return oidc.Token{}, fmt.Errorf(
			"AWS session expires too soon to mint an OIDC token (need at least %d seconds remaining)",
			minTokenDurationSeconds,
		)
	}

	var lastErr error
	for _, seconds := range durationAttempts(duration) {
		output, err := s.client.GetWebIdentityToken(ctx, &sts.GetWebIdentityTokenInput{
			Audience:         []string{s.audience},
			SigningAlgorithm: awssdk.String(signingAlgorithm),
			DurationSeconds:  awssdk.Int32(seconds),
		})
		if err != nil {
			lastErr = err
			if isSessionDurationEscalation(err) {
				continue
			}
			return oidc.Token{}, err
		}
		if output == nil {
			return oidc.Token{}, errors.New("STS returned an empty GetWebIdentityToken response")
		}

		jwt := awssdk.ToString(output.WebIdentityToken)
		if jwt == "" {
			return oidc.Token{}, errors.New("STS returned an empty web identity token")
		}

		expiry := time.Now().UTC().Add(time.Duration(seconds) * time.Second)
		if output.Expiration != nil {
			expiry = output.Expiration.UTC()
		}
		return oidc.Token{JWT: jwt, Expiry: expiry}, nil
	}
	return oidc.Token{}, lastErr
}

func durationSecondsForSession(credentials awssdk.Credentials, now time.Time) int32 {
	if !credentials.CanExpire {
		return tokenDurationSeconds
	}

	seconds := int64((credentials.Expires.Sub(now) - sessionExpirySafetySkew) / time.Second)
	if seconds >= int64(tokenDurationSeconds) {
		return tokenDurationSeconds
	}
	if seconds < int64(minTokenDurationSeconds) {
		return 0
	}
	return int32(seconds)
}

func durationAttempts(preferred int32) []int32 {
	if preferred <= minTokenDurationSeconds {
		return []int32{minTokenDurationSeconds}
	}
	return []int32{preferred, minTokenDurationSeconds}
}

func isSessionDurationEscalation(err error) bool {
	var apiErr smithy.APIError
	return errors.As(err, &apiErr) && apiErr.ErrorCode() == "SessionDurationEscalationException"
}
