package entities

type StatsResponse struct {
	AverageNight float64 `json:"avg_night"`
	MinimumNight float64 `json:"min_night"`
	MaximumNight float64 `json:"max_night"`
}

func FromStatsCalculation(calculation StatsCalculation) StatsResponse {
	return StatsResponse{
		AverageNight: calculation.AverageNight,
		MinimumNight: calculation.MinimumNight,
		MaximumNight: calculation.MaximumNight,
	}
}
