package useCases

import (
	"github.com/KKrusti/booking/domain"
	"github.com/KKrusti/booking/domain/valueobjects"
)

func CalculateStats(bookings []domain.Booking) valueobjects.Stats {
	return domain.CalcStats(bookings)
}
