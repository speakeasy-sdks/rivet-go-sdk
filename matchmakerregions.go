package rivet

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/rivet-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/rivet-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/rivet-go-sdk/pkg/utils"
	"net/http"
	"strings"
)

type matchmakerRegions struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newMatchmakerRegions(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *matchmakerRegions {
	return &matchmakerRegions{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// RegionsServiceList - Returns a list of regions available to this namespace.
//
// Regions are sorted by most optimal to least optimal. The player's IP address
// is used to calculate the regions' optimality.
func (s *matchmakerRegions) RegionsServiceList(ctx context.Context, request operations.RegionsServiceListRequest) (*operations.RegionsServiceListResponse, error) {
	baseURL := operations.RegionsServiceListServerList[0]
	if request.ServerURL != nil {
		baseURL = *request.ServerURL
	}

	url := strings.TrimSuffix(baseURL, "/") + "/regions"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := utils.ConfigureSecurityClient(s.defaultClient, request.Security)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.RegionsServiceListResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.MatchmakerListRegionsOutput
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.MatchmakerListRegionsOutput = out
		}
	}

	return res, nil
}
