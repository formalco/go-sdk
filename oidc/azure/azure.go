package azure

import (
	"context"
	"errors"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/samber/mo"

	"github.com/formalco/go-sdk/v3/oidc"
)

const armScope = "https://management.azure.com/.default"

type tokenSource struct {
	credential    azcore.TokenCredential
	integrationID string
}

// NewTokenSource returns a TokenSource backed by an Azure credential. It requests
// Microsoft Entra access tokens for Azure Resource Manager's fixed audience.
func NewTokenSource(credential azcore.TokenCredential, integrationID string) (oidc.TokenSource, error) {
	if credential == nil {
		return nil, errors.New("oidc: azure credential is required")
	}
	if err := oidc.ValidateAudience(oidc.AudiencePrefix + integrationID); err != nil {
		return nil, err
	}
	return &tokenSource{
		credential:    credential,
		integrationID: integrationID,
	}, nil
}

func (s *tokenSource) Token(ctx context.Context) (oidc.Token, error) {
	token, err := s.credential.GetToken(ctx, policy.TokenRequestOptions{
		Scopes: []string{armScope},
	})
	if err != nil {
		return oidc.Token{}, fmt.Errorf("get Azure Resource Manager token: %w", err)
	}
	if token.Token == "" {
		return oidc.Token{}, errors.New("azure credential returned an empty token")
	}
	return oidc.Token{
		JWT:                 token.Token,
		Expiry:              token.ExpiresOn,
		HeaderIntegrationID: mo.Some(s.integrationID),
	}, nil
}
