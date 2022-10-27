package entities

import (
	utils "github.com/KKrusti/booking/internal/core"
	"sort"
	"sync"
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
	return utils.Round(totalProfit)
}

func (booking Booking) CalcProfit() float64 {
	totalProfit := booking.CalcTotalProfit()
	profit := totalProfit / float64(booking.Nights)
	return utils.Round(profit)
}

func (booking Booking) GetCheckoutDate() time.Time {
	checkinDate := utils.StringToTime(booking.Checkin)
	checkoutDate := checkinDate.AddDate(0, 0, booking.Nights)
	return checkoutDate
}

// IsValidCombination checks whether a booking combination dates are compatible or if they overlap.
func IsValidCombination(bookings []Booking) bool {
	sortByCheckinDate(bookings)
	for i := 0; i < len(bookings)-1; i++ {
		currentCheckout := bookings[i].GetCheckoutDate()
		nextCheckin := utils.StringToTime(bookings[i+1].Checkin)
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
func GenerateAllCombinations(ch chan []Booking, wg *sync.WaitGroup, bookings []Booking) {
	length := len(bookings)
	defer close(ch)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		wg.Add(1)
		var subset []Booking
		for booking := 0; booking < length; booking++ {
			if (subsetBits>>booking)&1 == 1 {
				subset = append(subset, bookings[booking])
			}
		}
		ch <- subset
	}
}

func GetBooking(id string, checkin string, nights int, sellingRate, margin float64) Booking {
	return Booking{
		Id:          id,
		Checkin:     checkin,
		Nights:      nights,
		SellingRate: sellingRate,
		Margin:      margin,
	}
}
