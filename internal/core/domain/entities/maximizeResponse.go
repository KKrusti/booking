package entities

type MaxResponse struct {
	RequestIds    []string      `json:"request_ids"`
	TotalProfit   float64       `json:"total_profit"`
	StatsResponse StatsResponse `json:"stats_response"`
}
