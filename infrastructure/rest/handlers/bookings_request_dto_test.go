package handlers

import (
	"github.com/KKrusti/booking/domain"
	"github.com/KKrusti/booking/domain/valueobjects"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MapToDomain(t *testing.T) {
	dto := []BookingsRequestDTO{
		{
			Id:          "A",
			Checkin:     "2020-05-05",
			Nights:      10,
			SellingRate: 630,
			Margin:      15,
		}, {
			Id:          "B",
			Checkin:     "2020-06-07",
			Nights:      3,
			SellingRate: 150,
			Margin:      31,
		},
	}

	bookings := mapDtoToDomain(dto)

	expectedBookings := valueobjects.Bookings{Bookings: []domain.Booking{
		{
			Id:          "A",
			Checkin:     "2020-05-05",
			Nights:      10,
			SellingRate: 630,
			Margin:      15,
		},
		{
			Id:          "B",
			Checkin:     "2020-06-07",
			Nights:      3,
			SellingRate: 150,
			Margin:      31,
		},
	},
	}

	assert.Equal(t, expectedBookings, bookings)
}
