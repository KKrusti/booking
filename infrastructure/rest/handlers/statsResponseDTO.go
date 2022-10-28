package handlers

import "github.com/KKrusti/booking/domain/valueobjects"

type StatsResponseDTO struct {
	AverageNight float64 `json:"avg_night"`
	MinimumNight float64 `json:"min_night"`
	MaximumNight float64 `json:"max_night"`
}

func FromDomain(stats valueobjects.Stats) StatsResponseDTO {
	return StatsResponseDTO{
		MinimumNight: stats.MinimumNight,
		MaximumNight: stats.MaximumNight,
		AverageNight: stats.AverageNight,
	}
}
