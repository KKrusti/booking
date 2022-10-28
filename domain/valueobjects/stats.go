package valueobjects

type Stats struct {
	RequestIds   []string
	AverageNight float64
	MinimumNight float64
	MaximumNight float64
	TotalProfit  float64
}

func (stats Stats) SortByMostProfitable(comparison Stats) bool {
	return stats.TotalProfit > comparison.TotalProfit
}
