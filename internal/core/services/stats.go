package services

import (
	"github.com/KKrusti/booking/internal/core"
	"github.com/KKrusti/booking/internal/core/domain/entities"
)

func CalcStats(bookings []entities.Booking) entities.StatsCalculation {
	profit := make([]float64, len(bookings))
	profitPerNight := make([]float64, len(bookings))

	var requestIds []string
	minimum, maximum, totalprofit, averageNight := 0.0, 0.0, 0.0, 0.0
	for i := 0; i < len(bookings); i++ {
		requestIds = append(requestIds, bookings[i].Id)
		profit[i] = bookings[i].CalcTotalProfit()
		totalprofit += profit[i]
		profitPerNight[i] = profit[i] / float64(bookings[i].Nights)
		minimum = calcMinimum(i, minimum, profitPerNight[i])
		maximum = calcMaximum(maximum, profitPerNight[i])
	}
	averageNight = calcAverageNight(profitPerNight)

	return entities.StatsCalculation{
		RequestIds:   requestIds,
		AverageNight: averageNight,
		MinimumNight: minimum,
		MaximumNight: maximum,
		TotalProfit:  totalprofit,
	}
}

func calcAverageNight(profits []float64) float64 {
	sumProfits := 0.00
	for _, profit := range profits {
		sumProfits += profit
	}

	average := sumProfits / float64(len(profits))
	return utils.Round(average)
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
