package entities

//	type MaxResponse struct {
//		RequestIds    []string      `json:"request_ids"`
//		TotalProfit   float64       `json:"total_profit"`
//		StatsResponse StatsResponse `json:"stats_response"`
//	}
type MaxResponse struct {
	RequestIds   []string `json:"request_ids"`
	TotalProfit  float64  `json:"total_profit"`
	AverageNight float64  `json:"avg_night"`
	MinimumNight float64  `json:"min_night"`
	MaximumNight float64  `json:"max_night"`
}
