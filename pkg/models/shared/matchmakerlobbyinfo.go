package shared

// MatchmakerLobbyInfo
// A public lobby in the lobby list.
type MatchmakerLobbyInfo struct {
	GameModeID       string `json:"game_mode_id"`
	LobbyID          string `json:"lobby_id"`
	MaxPlayersDirect int64  `json:"max_players_direct"`
	MaxPlayersNormal int64  `json:"max_players_normal"`
	MaxPlayersParty  int64  `json:"max_players_party"`
	RegionID         string `json:"region_id"`
	TotalPlayerCount int64  `json:"total_player_count"`
}
