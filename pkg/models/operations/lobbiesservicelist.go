package operations

import (
	"github.com/speakeasy-sdks/rivet-go-sdk/pkg/models/shared"
)

var LobbiesServiceListServerList = []string{
	"https://party.api.rivet.gg/v1",
}

type LobbiesServiceListSecurity struct {
	BearerAuth *shared.SchemeBearerAuth `security:"scheme,type=http,subtype=bearer"`
}

type LobbiesServiceListRequest struct {
	Security  LobbiesServiceListSecurity
	ServerURL *string
}

type LobbiesServiceListResponse struct {
	ContentType                 string
	MatchmakerListLobbiesOutput *shared.MatchmakerListLobbiesOutput
	StatusCode                  int
}
