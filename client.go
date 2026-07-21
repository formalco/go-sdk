package formal

import (
	"errors"
	"net/http"
)

type options struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

// Option configures a Client.
type Option func(*options)

// WithAPIKey authenticates requests with an X-Api-Key header.
func WithAPIKey(apiKey string) Option {
	return func(o *options) {
		o.apiKey = apiKey
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
	if o.apiKey == "" {
		return nil, errors.New("formal: authentication required (use WithAPIKey)")
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
	client.Transport = &apiKeyTransport{
		apiKey: o.apiKey,
		base:   base,
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
