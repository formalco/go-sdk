package azure

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/samber/mo"

	"github.com/formalco/go-sdk/v3/oidc"
)

const armScope = "https://management.azure.com/.default"

type tokenSource struct {
	credential    azcore.TokenCredential
	integrationID string
}

// NewDefaultCredential returns Workload Identity when AZURE_FEDERATED_TOKEN_FILE
// is set, otherwise managed identity. AZURE_CLIENT_ID selects a user-assigned
// identity. This is not azidentity.NewDefaultAzureCredential.
func NewDefaultCredential() (azcore.TokenCredential, error) {
	if os.Getenv("AZURE_FEDERATED_TOKEN_FILE") != "" {
		return azidentity.NewWorkloadIdentityCredential(nil)
	}

	options := &azidentity.ManagedIdentityCredentialOptions{}
	if clientID := os.Getenv("AZURE_CLIENT_ID"); clientID != "" {
		options.ID = azidentity.ClientID(clientID)
	}
	return azidentity.NewManagedIdentityCredential(options)
}

// NewDefaultTokenSource builds a TokenSource from NewDefaultCredential.
func NewDefaultTokenSource(integrationID string) (oidc.TokenSource, error) {
	credential, err := NewDefaultCredential()
	if err != nil {
		return nil, fmt.Errorf("load Azure credential for OIDC: %w", err)
	}
	return NewTokenSource(credential, integrationID)
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
