package useCases

import (
	"github.com/KKrusti/booking/domain/valueobjects"
	"sort"
)

func Maximize(bookings valueobjects.Bookings) valueobjects.Stats {
	return processAllCombinations(bookings)
}

func processAllCombinations(bookings valueobjects.Bookings) valueobjects.Stats {
	validCombinations := generateCombinationsAndGetValid(bookings)
	allCalculations := calculateProfits(validCombinations)
	sortByMostProfitableBooking(allCalculations)
	return allCalculations[0]
}

func generateCombinationsAndGetValid(bookings valueobjects.Bookings) []valueobjects.Bookings {
	combinationsChan := make(chan valueobjects.Bookings)
	go bookings.GenerateAllCombinations(combinationsChan)

	return valueobjects.CheckValidCombinations(combinationsChan)
}

func sortByMostProfitableBooking(calculations []valueobjects.Stats) {
	sort.Slice(calculations[:], func(i, j int) bool {
		return calculations[i].IsMoreProfitableThan(calculations[j])
	})
}

func calculateProfits(combinations []valueobjects.Bookings) []valueobjects.Stats {
	var calculationsPerCombination []valueobjects.Stats
	for _, booking := range combinations {
		statsForBookings := booking.CalcStats()
		calculationsPerCombination = append(calculationsPerCombination, statsForBookings)
	}
	return calculationsPerCombination
}
