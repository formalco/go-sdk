package sdk

import "net/http"

type transport struct {
	underlyingTransport http.RoundTripper
	apiKey              string
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("X-Api-Key", t.apiKey)
	return t.underlyingTransport.RoundTrip(req)
}

func NewClient(apiKey string) *http.Client {
	return &http.Client{
		Transport: &transport{
			apiKey:              apiKey,
			underlyingTransport: http.DefaultTransport,
		},
	}
}
