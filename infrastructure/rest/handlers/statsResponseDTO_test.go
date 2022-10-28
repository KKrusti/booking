package handlers

import (
	"github.com/KKrusti/booking/domain/valueobjects"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MapFromDomain(t *testing.T) {
	domain := valueobjects.Stats{
		AverageNight: 15.2,
		MinimumNight: float64(8),
		MaximumNight: 13.2,
	}

	gotDto := FromDomain(domain)
	expectedDto := StatsResponseDTO{
		AverageNight: 15.2,
		MinimumNight: float64(8),
		MaximumNight: 13.2,
	}

	assert.Equal(t, expectedDto, gotDto)
}
