package shared

type CommonsMatchmakerLobbyJoinInfoPort struct {
	Host      *string                                 `json:"host,omitempty"`
	Hostname  string                                  `json:"hostname"`
	IsTLS     bool                                    `json:"is_tls"`
	Port      *int64                                  `json:"port,omitempty"`
	PortRange CommonsMatchmakerLobbyJoinInfoPortRange `json:"port_range"`
}
