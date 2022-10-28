package entities

import "github.com/KKrusti/booking/internal/core/domain/dto"

type StatsCalculation struct {
	RequestIds   []string
	AverageNight float64
	MinimumNight float64
	MaximumNight float64
	TotalProfit  float64
}

func (stats StatsCalculation) ToMaxResponse() dto.MaxResponse {
	return dto.MaxResponse{
		RequestIds:  stats.RequestIds,
		TotalProfit: stats.TotalProfit,
		StatsResponse: dto.StatsResponse{
			AverageNight: stats.AverageNight,
			MinimumNight: stats.MinimumNight,
			MaximumNight: stats.MaximumNight,
		},
	}
}

func (stats StatsCalculation) ToStatsResponse() dto.StatsResponse {
	return dto.StatsResponse{
		AverageNight: stats.AverageNight,
		MinimumNight: stats.MinimumNight,
		MaximumNight: stats.MaximumNight,
	}
}
