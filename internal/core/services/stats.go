package services

import "github.com/KKrusti/booking/internal/core/domain"

func CalcStats(requests domain.BookingRequest) (float64, float64, float64) {
	profit := make([]float64, len(requests.Requests))
	minimum, maximum := 0.0, 0.0
	for i, request := range requests.Requests {
		profit[i] = calcProfit(request)
		minimum = calcMinimum(i, minimum, profit[i])
		maximum = calcMaximum(maximum, profit[i])
	}

	averageNight := calcAverageNight(profit)
	return averageNight, minimum, maximum
}

func calcAverageNight(profits []float64) float64 {
	sumProfits := 0.00
	for _, profit := range profits {
		sumProfits += profit
	}
	return sumProfits / float64(len(profits))

}

func calcProfit(request domain.Request) float64 {
	totalProfit := (request.Margin / 100) * request.SellingRate
	return totalProfit / float64(request.Nights)
}

func calcMinimum(index int, minimum, profit float64) float64 {
	if index == 0 || minimum > profit {
		return profit
	} else {
		return minimum
	}
}

func calcMaximum(maximum, profit float64) float64 {
	if maximum > profit {
		return maximum
	} else {
		return profit
	}
}
