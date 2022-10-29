package useCases

import (
	"github.com/KKrusti/booking/domain/valueobjects"
)

func Maximize(bookings valueobjects.Bookings) valueobjects.Stats {
	return bookings.ProcessAllCombinations()
}
