package sdk

import "net/http"

type transport struct {
	underlyingTransport http.RoundTripper
	apiKey              string
	apiVersion          string
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("X-Api-Key", t.apiKey)
	req.Header.Add("X-Formal-API-Version", t.apiVersion)
	return t.underlyingTransport.RoundTrip(req)
}
