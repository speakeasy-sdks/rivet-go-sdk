package rivet

import (
	"github.com/speakeasy-sdks/rivet-go-sdk/pkg/utils"
	"net/http"
	"time"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Rivet struct {
	MatchmakerLobbies *matchmakerLobbies
	MatchmakerRegions *matchmakerRegions

	// Non-idiomatic field names below are to namespace fields from the fields names above to avoid name conflicts
	_defaultClient  HTTPClient
	_securityClient HTTPClient

	_serverURL  string
	_language   string
	_sdkVersion string
	_genVersion string
}

type SDKOption func(*Rivet)

func WithServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Rivet) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk._serverURL = serverURL
	}
}

func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Rivet) {
		sdk._defaultClient = client
	}
}

func New(opts ...SDKOption) *Rivet {
	sdk := &Rivet{
		_language:   "go",
		_sdkVersion: "0.8.0",
		_genVersion: "1.5.4",
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk._defaultClient == nil {
		sdk._defaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk._securityClient == nil {

		sdk._securityClient = sdk._defaultClient

	}

	sdk.MatchmakerLobbies = newMatchmakerLobbies(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.MatchmakerRegions = newMatchmakerRegions(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	return sdk
}
