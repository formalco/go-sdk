package formal

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/formalco/go-sdk/v3/oidc"
)

type options struct {
	baseURL         string
	apiKey          string
	oidcTokenSource oidc.TokenSource
	oidcSet         bool
	httpClient      *http.Client
}

// Option configures a Client.
type Option func(*options)

// WithAPIKey authenticates requests with an X-Api-Key header.
func WithAPIKey(apiKey string) Option {
	return func(o *options) {
		o.apiKey = apiKey
	}
}

// WithOIDCTokenSource authenticates requests with an Authorization Bearer JWT
// from source. Mutually exclusive with WithAPIKey.
func WithOIDCTokenSource(source oidc.TokenSource) Option {
	return func(o *options) {
		o.oidcTokenSource = source
		o.oidcSet = true
	}
}

// WithBaseURL sets the control-plane base URL. Defaults to DefaultURL.
func WithBaseURL(baseURL string) Option {
	return func(o *options) {
		o.baseURL = baseURL
	}
}

// WithHTTPClient uses a custom HTTP client. Its Transport is wrapped to attach
// Formal auth/version headers; Timeout and other client fields are preserved.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(o *options) {
		o.httpClient = httpClient
	}
}

// New constructs a Client from options.
func New(opts ...Option) (*Client, error) {
	o := options{
		baseURL: DefaultURL,
	}
	for _, opt := range opts {
		opt(&o)
	}

	if o.baseURL == "" {
		return nil, errors.New("formal: base URL must not be empty")
	}
	if o.apiKey != "" && o.oidcSet {
		return nil, errors.New("formal: WithAPIKey and WithOIDCTokenSource are mutually exclusive")
	}
	if o.oidcSet && o.oidcTokenSource == nil {
		return nil, errors.New("formal: WithOIDCTokenSource requires a TokenSource")
	}
	if o.apiKey == "" && !o.oidcSet {
		return nil, errors.New("formal: authentication required (use WithAPIKey or WithOIDCTokenSource)")
	}

	httpClient := o.httpClient
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	base := httpClient.Transport
	if base == nil {
		base = http.DefaultTransport
	}

	client := *httpClient
	if o.oidcSet {
		client.Transport = &oidcTransport{
			source: o.oidcTokenSource,
			base:   base,
		}
	} else {
		client.Transport = &apiKeyTransport{
			apiKey: o.apiKey,
			base:   base,
		}
	}

	return newClient(&client, o.baseURL), nil
}

type apiKeyTransport struct {
	apiKey string
	base   http.RoundTripper
}

func (t *apiKeyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req = req.Clone(req.Context())
	req.Header.Set("X-Api-Key", t.apiKey)
	req.Header.Set("X-Formal-API-Version", defaultAPIVersion)
	return t.base.RoundTrip(req)
}

type oidcTransport struct {
	source oidc.TokenSource
	base   http.RoundTripper
}

func (t *oidcTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	token, err := t.source.Token(req.Context())
	if err != nil {
		return nil, fmt.Errorf("formal: get OIDC token: %w", err)
	}
	if strings.TrimSpace(token.JWT) == "" {
		return nil, errors.New("formal: OIDC token must not be empty")
	}
	headerIntegrationID, sendIntegrationIDHeader := token.HeaderIntegrationID.Get()
	if sendIntegrationIDHeader {
		if err := oidc.ValidateAudience(oidc.AudiencePrefix + headerIntegrationID); err != nil {
			return nil, fmt.Errorf("formal: invalid OIDC header integration ID: %w", err)
		}
	}

	req = req.Clone(req.Context())
	req.Header.Set("Authorization", "Bearer "+token.JWT)
	req.Header.Set("X-Formal-API-Version", defaultAPIVersion)
	if sendIntegrationIDHeader {
		req.Header.Set("X-Formal-OIDC-Integration-Id", headerIntegrationID)
	}
	return t.base.RoundTrip(req)
}
