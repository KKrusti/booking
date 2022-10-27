package entities

type StatsCalculation struct {
	RequestIds   []string
	AverageNight float64
	MinimumNight float64
	MaximumNight float64
	TotalProfit  float64
}

func (stats StatsCalculation) ToMaxResponse() MaxResponse {
	return MaxResponse{
		RequestIds:   stats.RequestIds,
		TotalProfit:  stats.TotalProfit,
		AverageNight: stats.AverageNight,
		MinimumNight: stats.MinimumNight,
		MaximumNight: stats.MaximumNight,
	}
}
