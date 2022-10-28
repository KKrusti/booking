package entities

import (
	"github.com/KKrusti/booking/internal/core/domain/dto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ToMaxResponse(t *testing.T) {
	statsCalculation := StatsCalculation{
		RequestIds:   []string{"A", "B"},
		AverageNight: 20.5,
		MinimumNight: 7.5,
		MaximumNight: 8.9,
		TotalProfit:  45,
	}

	gotMaxResponseDto := statsCalculation.ToMaxResponse()

	expectedMaxResponseDTO := dto.MaxResponse{
		RequestIds:  []string{"A", "B"},
		TotalProfit: 45,
		StatsResponse: dto.StatsResponse{
			AverageNight: 20.5,
			MinimumNight: 7.5,
			MaximumNight: 8.9,
		},
	}
	assert.Equal(t, expectedMaxResponseDTO, gotMaxResponseDto)
}

func Test_ToStatsResponse(t *testing.T) {
	statsCalculation := StatsCalculation{
		RequestIds:   []string{"A", "B"},
		AverageNight: 20.5,
		MinimumNight: 7.5,
		MaximumNight: 8.9,
		TotalProfit:  45,
	}

	gotStatsResponse := statsCalculation.ToStatsResponse()
	expectedStatsResponseDTO := dto.StatsResponse{
		AverageNight: 20.5,
		MinimumNight: 7.5,
		MaximumNight: 8.9,
	}
	assert.Equal(t, expectedStatsResponseDTO, gotStatsResponse)
}
