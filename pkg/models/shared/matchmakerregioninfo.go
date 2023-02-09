package shared

// MatchmakerRegionInfo
// A region that the player can connect to.
type MatchmakerRegionInfo struct {
	DatacenterCoord              CommonsCoord    `json:"datacenter_coord"`
	DatacenterDistanceFromClient CommonsDistance `json:"datacenter_distance_from_client"`
	ProviderDisplayName          string          `json:"provider_display_name"`
	RegionDisplayName            string          `json:"region_display_name"`
	RegionID                     string          `json:"region_id"`
}
