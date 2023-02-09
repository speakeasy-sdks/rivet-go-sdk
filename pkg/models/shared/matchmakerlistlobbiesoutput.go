package shared

type MatchmakerListLobbiesOutput struct {
	GameModes []MatchmakerGameModeInfo `json:"game_modes"`
	Lobbies   []MatchmakerLobbyInfo    `json:"lobbies"`
	Regions   []MatchmakerRegionInfo   `json:"regions"`
}
