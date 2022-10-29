package handlers

import (
	"github.com/KKrusti/booking/domain"
	"github.com/KKrusti/booking/domain/valueobjects"
)

type BookingsRequestDTO struct {
	Id          string  `json:"request_id"`
	Checkin     string  `json:"check_in"`
	Nights      int     `json:"Nights"`
	SellingRate float64 `json:"selling_rate"`
	Margin      float64 `json:"Margin"`
}

func mapDtoToDomain(dto []BookingsRequestDTO) valueobjects.Bookings {
	var bookings valueobjects.Bookings
	for _, element := range dto {
		booking := domain.CreateBooking(element.Id, element.Checkin, element.Nights, element.SellingRate, element.Margin)
		bookings.Bookings = append(bookings.Bookings, booking)
	}
	return bookings
}
