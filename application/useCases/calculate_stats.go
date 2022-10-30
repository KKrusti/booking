package useCases

import (
	"github.com/KKrusti/booking/domain/valueobjects"
)

func CalculateStats(bookings valueobjects.Bookings) valueobjects.Stats {
	return bookings.CalcStats()
}
