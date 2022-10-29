package useCases

import (
	"github.com/KKrusti/booking/domain"
	"github.com/KKrusti/booking/domain/valueobjects"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Maximize(t *testing.T) {
	bookings := valueobjects.Bookings{
		Bookings: []domain.Booking{
			{
				Id:          "A",
				Checkin:     "2020-01-01",
				Nights:      5,
				SellingRate: 200,
				Margin:      20,
			},
			{
				Id:          "B",
				Checkin:     "2020-01-04",
				Nights:      4,
				SellingRate: 156,
				Margin:      5,
			},
			{
				Id:          "C",
				Checkin:     "2020-01-04",
				Nights:      4,
				SellingRate: 150,
				Margin:      6,
			},
			{
				Id:          "D",
				Checkin:     "2020-01-10",
				Nights:      4,
				SellingRate: 160,
				Margin:      30,
			},
		},
	}

	result := Maximize(bookings)

	expectedStats := valueobjects.Stats{
		RequestIds:   []string{"A", "D"},
		TotalProfit:  88,
		AverageNight: 10,
		MinimumNight: 8,
		MaximumNight: 12,
	}

	assert.Equal(t, expectedStats, result)

}
