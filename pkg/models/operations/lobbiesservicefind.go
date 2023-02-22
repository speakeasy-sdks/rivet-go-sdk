package operations

import (
	"github.com/speakeasy-sdks/rivet-go-sdk/pkg/models/shared"
)

var LobbiesServiceFindServerList = []string{
	"https://party.api.rivet.gg/v1",
}

type LobbiesServiceFindHeaders struct {
	Origin *string `header:"style=simple,explode=false,name=origin"`
}

type LobbiesServiceFindRequestBodyCaptcha1TypeEnum string

const (
	LobbiesServiceFindRequestBodyCaptcha1TypeEnumHcaptcha LobbiesServiceFindRequestBodyCaptcha1TypeEnum = "hcaptcha"
)

// LobbiesServiceFindRequestBodyCaptcha1
// Captcha configuration.
type LobbiesServiceFindRequestBodyCaptcha1 struct {
	ClientResponse string                                         `json:"client_response"`
	Type           *LobbiesServiceFindRequestBodyCaptcha1TypeEnum `json:"type,omitempty"`
}

type LobbiesServiceFindRequestBody struct {
	Captcha                *interface{} `json:"captcha,omitempty"`
	GameModes              []string     `json:"game_modes"`
	PreventAutoCreateLobby *bool        `json:"prevent_auto_create_lobby,omitempty"`
	Regions                []string     `json:"regions,omitempty"`
}

type LobbiesServiceFindSecurity struct {
	BearerAuth *shared.SchemeBearerAuth `security:"scheme,type=http,subtype=bearer"`
}

type LobbiesServiceFindRequest struct {
	Headers   LobbiesServiceFindHeaders
	Request   LobbiesServiceFindRequestBody `request:"mediaType=application/json"`
	Security  LobbiesServiceFindSecurity
	ServerURL *string
}

type LobbiesServiceFindResponse struct {
	ContentType               string
	MatchmakerFindLobbyOutput *shared.MatchmakerFindLobbyOutput
	StatusCode                int
}
