package domain

import (
	"github.com/KKrusti/booking/shared"
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
	profit := booking.GetProfitPerNight()
	if minimum >= profit {
		return shared.Round(profit)
	}
	return shared.Round(minimum)
}

func (booking Booking) CalcMaximum(maximum float64) float64 {
	profit := booking.GetProfitPerNight()
	if maximum <= profit {
		return shared.Round(profit)
	}

	return shared.Round(maximum)
}

func (booking Booking) GetProfitPerNight() float64 {
	return booking.CalcTotalProfit() / float64(booking.Nights)
}
