package valueobjects

type Stats struct {
	RequestIds   []string `json:"request_ids"`
	TotalProfit  float64  `json:"total_profit"`
	AverageNight float64  `json:"avg_night"`
	MinimumNight float64  `json:"min_night"`
	MaximumNight float64  `json:"max_night"`
}

func (stats Stats) IsMoreProfitableThan(comparison Stats) bool {
	return stats.TotalProfit > comparison.TotalProfit
}
