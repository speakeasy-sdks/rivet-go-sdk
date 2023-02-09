package operations

import (
	"github.com/speakeasy-sdks/rivet-go-sdk/pkg/models/shared"
)

var LobbiesServiceSetClosedServerList = []string{
	"https://party.api.rivet.gg/v1",
}

type LobbiesServiceSetClosedRequestBody struct {
	IsClosed bool `json:"is_closed"`
}

type LobbiesServiceSetClosedSecurity struct {
	BearerAuth shared.SchemeBearerAuth `security:"scheme,type=http,subtype=bearer"`
}

type LobbiesServiceSetClosedRequest struct {
	Request   LobbiesServiceSetClosedRequestBody `request:"mediaType=application/json"`
	Security  LobbiesServiceSetClosedSecurity
	ServerURL *string
}

type LobbiesServiceSetClosedResponse struct {
	ContentType string
	StatusCode  int64
}
