package operations

import (
	"github.com/speakeasy-sdks/rivet-go-sdk/pkg/models/shared"
)

var LobbiesServiceJoinServerList = []string{
	"https://party.api.rivet.gg/v1",
}

type LobbiesServiceJoinRequestBodyCaptcha1TypeEnum string

const (
	LobbiesServiceJoinRequestBodyCaptcha1TypeEnumHcaptcha LobbiesServiceJoinRequestBodyCaptcha1TypeEnum = "hcaptcha"
)

// LobbiesServiceJoinRequestBodyCaptcha1
// Captcha configuration.
type LobbiesServiceJoinRequestBodyCaptcha1 struct {
	ClientResponse string                                         `json:"client_response"`
	Type           *LobbiesServiceJoinRequestBodyCaptcha1TypeEnum `json:"type,omitempty"`
}

type LobbiesServiceJoinRequestBody struct {
	Captcha *interface{} `json:"captcha,omitempty"`
	LobbyID string       `json:"lobby_id"`
}

type LobbiesServiceJoinSecurity struct {
	BearerAuth *shared.SchemeBearerAuth `security:"scheme,type=http,subtype=bearer"`
}

type LobbiesServiceJoinRequest struct {
	Request   LobbiesServiceJoinRequestBody `request:"mediaType=application/json"`
	Security  LobbiesServiceJoinSecurity
	ServerURL *string
}

type LobbiesServiceJoinResponse struct {
	ContentType               string
	MatchmakerJoinLobbyOutput *shared.MatchmakerJoinLobbyOutput
	StatusCode                int
}
