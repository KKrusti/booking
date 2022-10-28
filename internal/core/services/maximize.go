package services

import (
	"github.com/KKrusti/booking/internal/core/domain/entities"
	"sort"
)

func Maximize(bookings []entities.Booking) entities.StatsCalculation {
	return processAllCombinations(bookings)
}

func processAllCombinations(bookings []entities.Booking) entities.StatsCalculation {
	validCombinations := generateCombinationsAndGetValid(bookings)
	allCalculations := calculateProfits(validCombinations)
	sortByMostProfitableBooking(allCalculations)
	return allCalculations[0]
}

func generateCombinationsAndGetValid(bookings []entities.Booking) [][]entities.Booking {
	combinationsChan := make(chan []entities.Booking)

	go entities.GenerateAllCombinations(combinationsChan, bookings)

	return entities.CheckValidCombinations(combinationsChan)
}

func sortByMostProfitableBooking(calculations []entities.StatsCalculation) {
	sort.Slice(calculations[:], func(i, j int) bool {
		return calculations[i].TotalProfit > calculations[j].TotalProfit
	})
}

func calculateProfits(combinations [][]entities.Booking) []entities.StatsCalculation {
	var calculationsPerCombination []entities.StatsCalculation
	for _, booking := range combinations {
		statsForBookings := CalcStats(booking)
		calculationsPerCombination = append(calculationsPerCombination, statsForBookings)
	}
	return calculationsPerCombination
}
