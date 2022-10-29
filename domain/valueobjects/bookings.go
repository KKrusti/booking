package valueobjects

import (
	"github.com/KKrusti/booking/domain"
	"github.com/KKrusti/booking/shared"
	"sort"
)

type Bookings struct {
	Bookings []domain.Booking
}

func (bookings Bookings) ProcessAllCombinations() Stats {
	combinationsChan := make(chan Bookings)
	go bookings.generateAllCombinations(combinationsChan)
	validCombinations := filterValidCombinations(combinationsChan)
	allCalculations := calculateProfits(validCombinations)
	sortByMostProfitableBooking(allCalculations)

	return allCalculations[0]
}

func (bookings Bookings) CalcStats() Stats {
	profitPerNight := make([]float64, len(bookings.Bookings))

	var requestIds []string
	minimum, maximum, totalProfit, averageNight := 0.0, 0.0, 0.0, 0.0
	for i := 0; i < len(bookings.Bookings); i++ {
		requestIds = append(requestIds, bookings.Bookings[i].Id)
		profitPerNight[i] = bookings.Bookings[i].GetProfitPerNight()
		totalProfit += bookings.Bookings[i].CalcTotalProfit()
		minimum = bookings.Bookings[i].CalcMinimum(minimum)
		if i == 0 {
			minimum = bookings.Bookings[i].GetProfitPerNight()
		}
		maximum = bookings.Bookings[i].CalcMaximum(maximum)
	}
	averageNight = calcAverageNight(profitPerNight)

	return Stats{
		RequestIds:   requestIds,
		AverageNight: averageNight,
		MinimumNight: minimum,
		MaximumNight: maximum,
		TotalProfit:  totalProfit,
	}
}

// sortByCheckinDate sorts all bookins from oldest to newest.
func (bookings Bookings) sortByCheckinDate() {
	sort.Slice(bookings.Bookings[:], func(i, j int) bool {
		return bookings.Bookings[i].Checkin < bookings.Bookings[j].Checkin
	})
}

// calculateProfits returns a slice with the stats of each combination.
func calculateProfits(combinations []Bookings) []Stats {
	var calculationsPerCombination []Stats
	for _, booking := range combinations {
		statsForBookings := booking.CalcStats()
		calculationsPerCombination = append(calculationsPerCombination, statsForBookings)
	}
	return calculationsPerCombination
}

// isValidBooking checks whether a booking combination dates are compatible or if they overlap.
func (bookings Bookings) isValidBooking() bool {
	bookings.sortByCheckinDate()
	for i := 0; i < len(bookings.Bookings)-1; i++ {
		currentCheckout := bookings.Bookings[i].GetCheckoutDate()
		nextCheckin := shared.StringToTime(bookings.Bookings[i+1].Checkin)
		if nextCheckin.Before(currentCheckout) {
			return false
		}
	}
	return true
}

// filterValidCombinations method that receives combinations through a channel and checks if they're valid or not. Only if it's
// a valid combination it's sent through another channel to process it.
func filterValidCombinations(combinations chan Bookings) []Bookings {
	var validCombinations []Bookings
	for combination := range combinations {
		if combination.isValidBooking() {
			validCombinations = append(validCombinations, combination)
		}
	}
	return validCombinations
}

// generateAllCombinations method that generates all combinations for given Bookings and sends each one through a channel
// to be processed as soon as it is generated.
func (bookings Bookings) generateAllCombinations(ch chan Bookings) {
	defer close(ch)
	length := len(bookings.Bookings)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset Bookings
		for index := 0; index < length; index++ {
			if (subsetBits>>index)&1 == 1 {
				subset.Bookings = append(subset.Bookings, bookings.Bookings[index])
			}
		}
		ch <- subset
	}
}

func calcAverageNight(profits []float64) float64 {
	sumProfits := 0.00
	for _, profit := range profits {
		sumProfits += profit
	}

	average := sumProfits / float64(len(profits))
	return shared.Round(average)
}
