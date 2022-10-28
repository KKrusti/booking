package domain

import (
	"github.com/KKrusti/booking/domain/valueobjects"
	"github.com/KKrusti/booking/shared"
	"sort"
	"time"
)

type Booking struct {
	Id          string  `json:"request_id"`
	Checkin     string  `json:"check_in"`
	Nights      int     `json:"nights"`
	SellingRate float64 `json:"selling_rate"`
	Margin      float64 `json:"margin"`
}

func (booking Booking) CalcTotalProfit() float64 {
	totalProfit := (booking.Margin / 100) * booking.SellingRate
	return shared.Round(totalProfit)
}

func (booking Booking) CalcProfit() float64 {
	totalProfit := booking.CalcTotalProfit()
	profit := totalProfit / float64(booking.Nights)
	return shared.Round(profit)
}

func (booking Booking) GetCheckoutDate() time.Time {
	checkinDate := shared.StringToTime(booking.Checkin)
	checkoutDate := checkinDate.AddDate(0, 0, booking.Nights)
	return checkoutDate
}

// IsValidBooking checks whether a booking combination dates are compatible or if they overlap.
func IsValidBooking(bookings []Booking) bool {
	sortByCheckinDate(bookings)
	for i := 0; i < len(bookings)-1; i++ {
		currentCheckout := bookings[i].GetCheckoutDate()
		nextCheckin := shared.StringToTime(bookings[i+1].Checkin)
		if nextCheckin.Before(currentCheckout) {
			return false
		}
	}
	return true
}

// sortByCheckinDate sorts all bookins from oldest to newest.
func sortByCheckinDate(requests []Booking) {
	sort.Slice(requests[:], func(i, j int) bool {
		return requests[i].Checkin < requests[j].Checkin
	})
}

// GenerateAllCombinations method that generates all combinations for given Bookings and sends each one through a channel
// to be processed as soon as it is generated.
func GenerateAllCombinations(ch chan []Booking, bookings []Booking) {
	defer close(ch)
	length := len(bookings)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset []Booking
		for booking := 0; booking < length; booking++ {
			if (subsetBits>>booking)&1 == 1 {
				subset = append(subset, bookings[booking])
			}
		}
		ch <- subset
	}
}

// CheckValidCombinations method that receives combinations through a channel and checks if they're valid or not. Only if it's
// a valid combination it's sent through another channel to process it.
func CheckValidCombinations(combinations chan []Booking) [][]Booking {
	var validCombinations [][]Booking
	for combination := range combinations {
		if IsValidBooking(combination) {
			validCombinations = append(validCombinations, combination)
		}
	}
	return validCombinations
}

func CreateBooking(id string, checkin string, nights int, sellingRate, margin float64) Booking {
	return Booking{
		Id:          id,
		Checkin:     checkin,
		Nights:      nights,
		SellingRate: sellingRate,
		Margin:      margin,
	}
}

func (booking Booking) CalcMinimum(minimum float64) float64 {
	profit := booking.GetProfit()
	if minimum >= profit {
		return shared.Round(profit)
	}
	return shared.Round(minimum)
}

func (booking Booking) CalcMaximum(maximum float64) float64 {
	profit := booking.GetProfit()
	if maximum <= profit {
		return shared.Round(profit)
	}

	return shared.Round(maximum)
}

func (booking Booking) GetProfit() float64 {
	return booking.CalcTotalProfit() / float64(booking.Nights)
}

func CalcStats(bookings []Booking) valueobjects.Stats {
	profitPerNight := make([]float64, len(bookings))

	var requestIds []string
	minimum, maximum, totalProfit, averageNight := 0.0, 0.0, 0.0, 0.0
	for i := 0; i < len(bookings); i++ {
		requestIds = append(requestIds, bookings[i].Id)
		profitPerNight[i] = bookings[i].GetProfit()
		minimum = bookings[i].CalcMinimum(minimum)
		if i == 0 {
			minimum = bookings[i].GetProfit()
		}
		maximum = bookings[i].CalcMaximum(maximum)
	}
	averageNight = calcAverageNight(profitPerNight)

	return valueobjects.Stats{
		RequestIds:   requestIds,
		AverageNight: averageNight,
		MinimumNight: minimum,
		MaximumNight: maximum,
		TotalProfit:  totalProfit,
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
