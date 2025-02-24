package sdk

import "net/http"

type transport struct {
	underlyingTransport http.RoundTripper
	apiKey              string
	apiVersion          string
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("X-Api-Key", t.apiKey)
	req.Header.Add("X-Formal-API-Version", "2025-02-24")
	return t.underlyingTransport.RoundTrip(req)
}

func NewClient(apiKey, apiVersion string) *http.Client {
	return &http.Client{
		Transport: &transport{
			apiKey:              apiKey,
			apiVersion:          apiVersion,
			underlyingTransport: http.DefaultTransport,
		},
	}
}
