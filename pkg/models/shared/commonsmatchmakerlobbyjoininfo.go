package shared

// CommonsMatchmakerLobbyJoinInfo
// A matchmaker lobby.
type CommonsMatchmakerLobbyJoinInfo struct {
	LobbyID string                                        `json:"lobby_id"`
	Player  CommonsMatchmakerLobbyJoinInfoPlayer          `json:"player"`
	Ports   map[string]CommonsMatchmakerLobbyJoinInfoPort `json:"ports"`
	Region  CommonsMatchmakerLobbyJoinInfoRegion          `json:"region"`
}
