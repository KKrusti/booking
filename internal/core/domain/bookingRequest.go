package domain

import "time"

type BookingRequest struct {
	Requests []Request `json:"requests"`
}

type Request struct {
	Id          string    `json:"request_id"`
	Checkin     time.Time `json:"check_in"`
	Nights      int       `json:"nights"`
	SellingRate float64   `json:"selling_rate"`
	Margin      float64   `json:"margin"`
}
