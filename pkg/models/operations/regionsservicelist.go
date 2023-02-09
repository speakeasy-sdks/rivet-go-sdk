package operations

import (
	"github.com/speakeasy-sdks/rivet-go-sdk/pkg/models/shared"
)

var RegionsServiceListServerList = []string{
	"https://party.api.rivet.gg/v1",
}

type RegionsServiceListSecurity struct {
	BearerAuth *shared.SchemeBearerAuth `security:"scheme,type=http,subtype=bearer"`
}

type RegionsServiceListRequest struct {
	Security  RegionsServiceListSecurity
	ServerURL *string
}

type RegionsServiceListResponse struct {
	ContentType                 string
	MatchmakerListRegionsOutput *shared.MatchmakerListRegionsOutput
	StatusCode                  int64
}
