package services

import (
	"github.com/KKrusti/booking/internal/core/domain/entities"
	"sort"
	"sync"
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
	wg := &sync.WaitGroup{}

	go entities.GenerateAllCombinations(combinationsChan, wg, bookings)

	return checkValidCombinations(combinationsChan, wg)
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

// checkValidCombinations method that receives combinations through a channel and checks if they're valid or not. Only if it's
// a valid combination it's sent through another channel to process it.
func checkValidCombinations(combinations chan []entities.Booking, wg *sync.WaitGroup) [][]entities.Booking {
	var validCombinations [][]entities.Booking
	for combination := range combinations {
		if entities.IsValidCombination(combination) {
			validCombinations = append(validCombinations, combination)
		}
		wg.Done()
	}
	wg.Wait()
	return validCombinations
}
