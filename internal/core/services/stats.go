package services

import (
	"github.com/KKrusti/booking/internal/core"
	"github.com/KKrusti/booking/internal/core/domain/entities"
)

func CalcStats(requests []entities.Request) (float64, float64, float64) {
	profit := make([]float64, len(requests))
	minimum, maximum := 0.0, 0.0
	for i := 0; i < len(requests); i++ {
		profit[i] = calcProfit(requests[i])
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

	average := sumProfits / float64(len(profits))
	return utils.Round(average)

}

func calcProfit(request entities.Request) float64 {
	totalProfit := (request.Margin / 100) * request.SellingRate
	profit := totalProfit / float64(request.Nights)
	return utils.Round(profit)
}

func calcMinimum(index int, minimum, profit float64) float64 {
	if index == 0 || minimum > profit {
		return utils.Round(profit)
	} else {
		return utils.Round(minimum)
	}
}

func calcMaximum(maximum, profit float64) float64 {
	if maximum > profit {
		return maximum
	} else {
		return profit
	}
}
