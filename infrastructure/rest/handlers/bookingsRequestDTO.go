package handlers

import (
	"github.com/KKrusti/booking/domain"
	"github.com/KKrusti/booking/domain/valueobjects"
)

type BookingsRequestDTO struct {
	Id          string  `validate:"required" json:"request_id"`
	Checkin     string  `validate:"required,len=10" json:"check_in"`
	Nights      int     `validate:"required,min=1" json:"nights"`
	SellingRate float64 `validate:"required,min=1" json:"selling_rate"`
	Margin      float64 `validate:"required,min=0" json:"margin"`
}

func mapDtoToDomain(dto []BookingsRequestDTO) valueobjects.Bookings {
	var bookings valueobjects.Bookings
	for _, element := range dto {
		booking := domain.CreateBooking(element.Id, element.Checkin, element.Nights, element.SellingRate, element.Margin)
		bookings.Bookings = append(bookings.Bookings, booking)
	}
	return bookings
}
