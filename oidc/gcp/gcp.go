// Package gcp provides an OIDC token source backed by ambient Google Cloud credentials.
package gcp

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"cloud.google.com/go/auth"
	"cloud.google.com/go/auth/credentials"
	"cloud.google.com/go/auth/credentials/idtoken"
	"cloud.google.com/go/auth/credentials/impersonate"
	"cloud.google.com/go/auth/httptransport"
	"github.com/samber/lo"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"

	"github.com/formalco/go-sdk/v3/oidc"
)

const (
	cloudPlatformScope          = "https://www.googleapis.com/auth/cloud-platform"
	applicationCredentialsEnv   = "GOOGLE_APPLICATION_CREDENTIALS"
	applicationCredentialsFile  = "application_default_credentials.json"
	impersonationURLField       = "service_account_impersonation_url"
	credentialsTypeExternal     = "external_account"
	credentialsTypeImpersonated = "impersonated_service_account"
)

type tokenSource struct {
	provider auth.TokenProvider
}

// NewDefaultTokenSource returns a token source that uses Google Application
// Default Credentials to mint Google-signed ID tokens for audience.
//
// On GCE, GKE, Cloud Run, and other Google Cloud runtimes, tokens come from
// the metadata server for the attached service account. Elsewhere, service
// account keys, impersonated service account configs, and Workload Identity
// Federation configs with service account impersonation are supported.
func NewDefaultTokenSource(audience string) (oidc.TokenSource, error) {
	return NewDefaultTokenSourceWithHTTPClient(audience, nil)
}

// NewDefaultTokenSourceWithHTTPClient is like NewDefaultTokenSource but uses
// httpClient as the base client for requests to Google.
func NewDefaultTokenSourceWithHTTPClient(audience string, httpClient *http.Client) (oidc.TokenSource, error) {
	if err := oidc.ValidateAudience(audience); err != nil {
		return nil, err
	}
	adc, err := readApplicationCredentialsFile()
	if err != nil {
		return nil, fmt.Errorf("read GCP application default credentials: %w", err)
	}
	var provider auth.TokenProvider
	switch gjson.GetBytes(adc, "type").String() {
	case credentialsTypeExternal, credentialsTypeImpersonated:
		provider, err = newImpersonatedIDTokenProvider(audience, adc, httpClient)
	default:
		provider, err = idtoken.NewCredentials(&idtoken.Options{
			Audience:           audience,
			ComputeTokenFormat: idtoken.ComputeTokenFormatFull,
			Client:             httpClient,
		})
	}
	if err != nil {
		return nil, fmt.Errorf("load GCP credentials for OIDC: %w", err)
	}
	return NewTokenSource(provider)
}

// NewTokenSource adapts a Google auth token provider that returns ID tokens.
func NewTokenSource(provider auth.TokenProvider) (oidc.TokenSource, error) {
	if provider == nil {
		return nil, errors.New("oidc: GCP token provider is required")
	}
	return &tokenSource{provider: provider}, nil
}

func (s *tokenSource) Token(ctx context.Context) (oidc.Token, error) {
	token, err := s.provider.Token(ctx)
	if err != nil {
		return oidc.Token{}, fmt.Errorf("get GCP ID token: %w", err)
	}
	if token == nil || token.Value == "" {
		return oidc.Token{}, errors.New("GCP credentials returned an empty ID token")
	}
	return oidc.Token{JWT: token.Value, Expiry: token.Expiry}, nil
}

// newImpersonatedIDTokenProvider mints ID tokens by calling IAM
// generateIdToken with the config's source credentials. idtoken.NewCredentials
// (as of cloud.google.com/go/auth v0.24.0) sends that request unauthenticated
// for these credential types: https://github.com/googleapis/google-cloud-go/issues/19939.
func newImpersonatedIDTokenProvider(audience string, adc []byte, httpClient *http.Client) (auth.TokenProvider, error) {
	impersonationURL := gjson.GetBytes(adc, impersonationURLField).String()
	if impersonationURL == "" {
		return nil, errors.New("GCP credentials config must set service_account_impersonation_url to mint ID tokens")
	}
	targetPrincipal, _, _ := strings.Cut(path.Base(impersonationURL), ":")

	var sourceJSON []byte
	if gjson.GetBytes(adc, "type").String() == credentialsTypeExternal {
		var err error
		sourceJSON, err = sjson.DeleteBytes(adc, impersonationURLField)
		if err != nil {
			return nil, err
		}
	} else {
		sourceJSON = []byte(gjson.GetBytes(adc, "source_credentials").Raw)
	}
	sourceCredentials, err := credentials.DetectDefault(&credentials.DetectOptions{
		CredentialsJSON:  sourceJSON,
		Scopes:           []string{cloudPlatformScope},
		Client:           httpClient,
		UseSelfSignedJWT: true,
	})
	if err != nil {
		return nil, fmt.Errorf("load source credentials: %w", err)
	}

	transportOptions := &httptransport.Options{Credentials: sourceCredentials}
	if httpClient != nil {
		transportOptions.BaseRoundTripper = httpClient.Transport
	}
	iamClient, err := httptransport.NewClient(transportOptions)
	if err != nil {
		return nil, err
	}
	delegates := lo.Map(gjson.GetBytes(adc, "delegates").Array(), func(delegate gjson.Result, _ int) string {
		return delegate.String()
	})
	return impersonate.NewIDTokenCredentials(&impersonate.IDTokenOptions{
		Audience:        audience,
		TargetPrincipal: targetPrincipal,
		Delegates:       delegates,
		IncludeEmail:    true,
		Client:          iamClient,
		Credentials:     sourceCredentials,
	})
}

func readApplicationCredentialsFile() ([]byte, error) {
	if filename := os.Getenv(applicationCredentialsEnv); filename != "" {
		return os.ReadFile(filename)
	}
	adc, err := os.ReadFile(wellKnownApplicationCredentialsFile())
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}
	return adc, err
}

func wellKnownApplicationCredentialsFile() string {
	if runtime.GOOS == "windows" {
		return filepath.Join(os.Getenv("APPDATA"), "gcloud", applicationCredentialsFile)
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "gcloud", applicationCredentialsFile)
}
