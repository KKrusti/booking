package valueobjects

import "sort"

type Stats struct {
	RequestIds   []string `json:"request_ids"`
	TotalProfit  float64  `json:"total_profit"`
	AverageNight float64  `json:"avg_night"`
	MinimumNight float64  `json:"min_night"`
	MaximumNight float64  `json:"max_night"`
}

func (stats Stats) isMoreProfitableThan(comparison Stats) bool {
	return stats.TotalProfit > comparison.TotalProfit
}

func sortByMostProfitableBooking(calculations []Stats) {
	sort.Slice(calculations[:], func(i, j int) bool {
		return calculations[i].isMoreProfitableThan(calculations[j])
	})
}
