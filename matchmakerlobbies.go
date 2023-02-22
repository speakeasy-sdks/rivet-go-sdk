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

type matchmakerLobbies struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newMatchmakerLobbies(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *matchmakerLobbies {
	return &matchmakerLobbies{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// LobbiesServiceFind - Finds a lobby based on the given criteria.
// If a lobby is not found and `prevent_auto_create_lobby` is `true`,
// a new lobby will be created.
func (s *matchmakerLobbies) LobbiesServiceFind(ctx context.Context, request operations.LobbiesServiceFindRequest) (*operations.LobbiesServiceFindResponse, error) {
	baseURL := operations.LobbiesServiceFindServerList[0]
	if request.ServerURL != nil {
		baseURL = *request.ServerURL
	}

	url := strings.TrimSuffix(baseURL, "/") + "/lobbies/find"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	utils.PopulateHeaders(ctx, req, request.Headers)

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

	res := &operations.LobbiesServiceFindResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.MatchmakerFindLobbyOutput
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.MatchmakerFindLobbyOutput = out
		}
	}

	return res, nil
}

// LobbiesServiceJoin - Joins a specific lobby.
// This request will use the direct player count configured for the
// lobby group.
func (s *matchmakerLobbies) LobbiesServiceJoin(ctx context.Context, request operations.LobbiesServiceJoinRequest) (*operations.LobbiesServiceJoinResponse, error) {
	baseURL := operations.LobbiesServiceJoinServerList[0]
	if request.ServerURL != nil {
		baseURL = *request.ServerURL
	}

	url := strings.TrimSuffix(baseURL, "/") + "/lobbies/join"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

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

	res := &operations.LobbiesServiceJoinResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.MatchmakerJoinLobbyOutput
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.MatchmakerJoinLobbyOutput = out
		}
	}

	return res, nil
}

// LobbiesServiceList - Lists all open lobbies.
func (s *matchmakerLobbies) LobbiesServiceList(ctx context.Context, request operations.LobbiesServiceListRequest) (*operations.LobbiesServiceListResponse, error) {
	baseURL := operations.LobbiesServiceListServerList[0]
	if request.ServerURL != nil {
		baseURL = *request.ServerURL
	}

	url := strings.TrimSuffix(baseURL, "/") + "/lobbies/list"

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

	res := &operations.LobbiesServiceListResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.MatchmakerListLobbiesOutput
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.MatchmakerListLobbiesOutput = out
		}
	}

	return res, nil
}

// LobbiesServiceSetClosed - If `is_closed` is `true`, players will be prevented from joining the lobby.
// Does not shutdown the lobby.
func (s *matchmakerLobbies) LobbiesServiceSetClosed(ctx context.Context, request operations.LobbiesServiceSetClosedRequest) (*operations.LobbiesServiceSetClosedResponse, error) {
	baseURL := operations.LobbiesServiceSetClosedServerList[0]
	if request.ServerURL != nil {
		baseURL = *request.ServerURL
	}

	url := strings.TrimSuffix(baseURL, "/") + "/lobbies/closed"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

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

	res := &operations.LobbiesServiceSetClosedResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 204:
	}

	return res, nil
}
