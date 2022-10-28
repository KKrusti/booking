package useCases

import (
	"github.com/KKrusti/booking/domain"
	"github.com/KKrusti/booking/domain/valueobjects"
	"sort"
)

func Maximize(bookings []domain.Booking) valueobjects.Stats {
	return processAllCombinations(bookings)
}

func processAllCombinations(bookings []domain.Booking) valueobjects.Stats {
	validCombinations := generateCombinationsAndGetValid(bookings)
	allCalculations := calculateProfits(validCombinations)
	sortByMostProfitableBooking(allCalculations)
	return allCalculations[0]
}

func generateCombinationsAndGetValid(bookings []domain.Booking) [][]domain.Booking {
	combinationsChan := make(chan []domain.Booking)
	go domain.GenerateAllCombinations(combinationsChan, bookings)

	return domain.CheckValidCombinations(combinationsChan)
}

func sortByMostProfitableBooking(calculations []valueobjects.Stats) {
	sort.Slice(calculations[:], func(i, j int) bool {
		return calculations[i].SortByMostProfitable(calculations[j])
	})
}

func calculateProfits(combinations [][]domain.Booking) []valueobjects.Stats {
	var calculationsPerCombination []valueobjects.Stats
	for _, booking := range combinations {
		statsForBookings := domain.CalcStats(booking)
		calculationsPerCombination = append(calculationsPerCombination, statsForBookings)
	}
	return calculationsPerCombination
}
