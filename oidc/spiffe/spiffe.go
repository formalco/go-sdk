// Package spiffe provides an OIDC token source backed by the SPIFFE Workload API.
package spiffe

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/spiffe/go-spiffe/v2/svid/jwtsvid"
	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/formalco/go-sdk/v3/oidc"
)

const refreshSkew = 30 * time.Second

type tokenSource struct {
	audience string
	addr     string

	mu     sync.Mutex
	cached oidc.Token
}

// NewTokenSource returns a TokenSource that fetches JWT-SVIDs for audience
// from the SPIFFE Workload API at SPIFFE_ENDPOINT_SOCKET.
//
// Formal verifies JWT-SVIDs through an OIDC integration that trusts the SPIRE
// OIDC Discovery Provider, so the SPIRE server must set jwt_issuer to the
// integration issuer.
func NewTokenSource(audience string) (oidc.TokenSource, error) {
	addr, _ := workloadapi.GetDefaultAddress()
	if addr == "" {
		return nil, fmt.Errorf("oidc: SPIFFE Workload API address is required in %s", workloadapi.SocketEnv)
	}
	if err := oidc.ValidateAudience(audience); err != nil {
		return nil, err
	}
	return &tokenSource{audience: audience, addr: addr}, nil
}

func (s *tokenSource) Token(ctx context.Context) (oidc.Token, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.cached.JWT != "" && time.Now().Add(refreshSkew).Before(s.cached.Expiry) {
		return s.cached, nil
	}

	svid, err := workloadapi.FetchJWTSVID(ctx, jwtsvid.Params{Audience: s.audience}, workloadapi.WithAddr(s.addr))
	if err != nil {
		if s.cached.JWT != "" && time.Now().Before(s.cached.Expiry) {
			return s.cached, nil
		}
		return oidc.Token{}, fmt.Errorf("fetch SPIFFE JWT-SVID: %w", err)
	}
	s.cached = oidc.Token{JWT: svid.Marshal(), Expiry: svid.Expiry}
	return s.cached, nil
}
