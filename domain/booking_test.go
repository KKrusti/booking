package domain

import (
	utils "github.com/KKrusti/booking/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_calcProfit(t *testing.T) {
	booking := Booking{
		Id:          "test",
		Nights:      5,
		SellingRate: 850,
		Margin:      17,
	}

	profit := booking.CalcProfit()
	expectedProfit := 28.9

	assert.Equal(t, expectedProfit, profit)
}

func Test_checkoutDate(t *testing.T) {
	booking := Booking{
		Nights:  5,
		Checkin: "2018-05-01",
	}

	checkoutDate := booking.GetCheckoutDate()
	expectedDate := utils.StringToTime("2018-05-06")

	assert.Equal(t, expectedDate, checkoutDate)
}

func Test_booking_getter(t *testing.T) {
	id := "A"
	checkin := "2020-05-05"
	nights := 6
	sellingRate := float64(400)
	margin := float64(13)

	gotBooking := CreateBooking(id, checkin, nights, sellingRate, margin)

	expectedBooking := Booking{
		Id:          id,
		Checkin:     checkin,
		Nights:      nights,
		SellingRate: sellingRate,
		Margin:      margin,
	}

	assert.Equal(t, expectedBooking, gotBooking)
}

func Test_CalcMinimum(t *testing.T) {
	type args struct {
		booking Booking
		value   float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "previous smaller",
			args: args{
				booking: Booking{
					Id:          "a",
					Checkin:     "2020-01-01",
					Nights:      5,
					SellingRate: 200,
					Margin:      20,
				},
				value: float64(10),
			},
			want: float64(8),
		},
		{
			name: "current smaller",
			args: args{
				booking: Booking{
					Id:          "a",
					Checkin:     "2020-01-01",
					Nights:      5,
					SellingRate: 200,
					Margin:      20,
				},
				value: float64(3),
			},
			want: float64(3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.booking.CalcMinimum(tt.args.value)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_CalcMaximum(t *testing.T) {
	type args struct {
		booking Booking
		value   float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "current bigger",
			args: args{
				booking: Booking{
					Id:          "a",
					Checkin:     "2020-01-01",
					Nights:      5,
					SellingRate: 200,
					Margin:      20,
				},
				value: float64(10),
			},
			want: float64(10),
		},
		{
			name: "previous bigger",
			args: args{
				booking: Booking{
					Id:          "a",
					Checkin:     "2020-01-01",
					Nights:      5,
					SellingRate: 200,
					Margin:      20,
				},
				value: float64(5),
			},
			want: float64(8),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.booking.CalcMaximum(tt.args.value)
			assert.Equal(t, tt.want, got)
		})
	}
}
