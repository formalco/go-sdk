package formal

import (
	_ "embed"
	"strings"
)

const (
	DefaultURL = "https://api.formal.ai"

	// defaultAPIVersion is the X-Formal-API-Version header value shipped with
	// this SDK release. Not overridable — bump when the SDK adopts a new API.
	defaultAPIVersion = "2026-01-29"
)

//go:embed VERSION
var version string

// userAgent is prepended to the connect-go User-Agent so the API can tell SDK
// callers apart from other Go clients.
var userAgent = "formal-go-sdk/" + strings.TrimSpace(version)
