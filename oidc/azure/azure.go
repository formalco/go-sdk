package azure

import (
	"context"
	"errors"
	"fmt"
	"net/http"
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
	return newDefaultCredential(nil)
}

func newDefaultCredential(httpClient *http.Client) (azcore.TokenCredential, error) {
	clientOptions := azcore.ClientOptions{}
	if httpClient != nil {
		clientOptions.Transport = httpClient
	}
	if os.Getenv("AZURE_FEDERATED_TOKEN_FILE") != "" {
		return azidentity.NewWorkloadIdentityCredential(&azidentity.WorkloadIdentityCredentialOptions{
			ClientOptions: clientOptions,
		})
	}

	options := &azidentity.ManagedIdentityCredentialOptions{ClientOptions: clientOptions}
	if clientID := os.Getenv("AZURE_CLIENT_ID"); clientID != "" {
		options.ID = azidentity.ClientID(clientID)
	}
	return azidentity.NewManagedIdentityCredential(options)
}

// NewDefaultTokenSource builds a TokenSource from NewDefaultCredential.
func NewDefaultTokenSource(integrationID string) (oidc.TokenSource, error) {
	return NewDefaultTokenSourceWithHTTPClient(integrationID, nil)
}

// NewDefaultTokenSourceWithHTTPClient builds a default Azure TokenSource whose
// identity-provider requests use httpClient.
func NewDefaultTokenSourceWithHTTPClient(integrationID string, httpClient *http.Client) (oidc.TokenSource, error) {
	credential, err := newDefaultCredential(httpClient)
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
