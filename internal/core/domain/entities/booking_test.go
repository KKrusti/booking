package entities

import (
	utils "github.com/KKrusti/booking/internal/core"
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

func Test_allCombinations(t *testing.T) {

	request := []Booking{
		{
			Id: "A",
		},
		{
			Id: "B",
		},
		{
			Id: "C",
		},
	}

	ch := make(chan []Booking)
	go GenerateAllCombinations(ch, request)

	var combination [][]Booking
	for received := range ch {
		combination = append(combination, received)
	}

	assert.Equal(t, 7, len(combination))

	a := []Booking{{Id: "A"}}
	b := []Booking{{Id: "B"}}
	c := []Booking{{Id: "C"}}
	ab := []Booking{{Id: "A"}, {Id: "B"}}
	bc := []Booking{{Id: "B"}, {Id: "C"}}
	ac := []Booking{{Id: "A"}, {Id: "C"}}
	abc := []Booking{{Id: "A"}, {Id: "B"}, {Id: "C"}}
	assert.Contains(t, combination, a)
	assert.Contains(t, combination, b)
	assert.Contains(t, combination, c)
	assert.Contains(t, combination, ab)
	assert.Contains(t, combination, ac)
	assert.Contains(t, combination, bc)
	assert.Contains(t, combination, abc)
}

func Test_validBooking(t *testing.T) {
	type args struct {
		request []Booking
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "large distance",
			args: args{
				request: []Booking{
					{
						Checkin: "2018-01-01",
						Nights:  5,
					},
					{
						Checkin: "2018-01-10",
						Nights:  5,
					},
				},
			},
			want: true,
		},
		{
			name: "finish and start same day",
			args: args{
				request: []Booking{
					{
						Checkin: "2018-01-01",
						Nights:  5,
					},
					{
						Checkin: "2018-01-06",
						Nights:  10,
					},
				},
			},
			want: true,
		},
		{
			name: "overlap",
			args: args{
				request: []Booking{
					{
						Checkin: "2018-01-01",
						Nights:  5,
					},
					{
						Checkin: "2018-01-04",
						Nights:  15,
					},
				},
			},
			want: false,
		},
		{
			name: "unsorted but valid",
			args: args{
				request: []Booking{
					{
						Checkin: "2018-01-01",
						Nights:  5,
					},
					{
						Checkin: "2018-01-20",
						Nights:  5,
					},
					{
						Checkin: "2018-01-10",
						Nights:  5,
					},
				},
			},
			want: true,
		},
		{
			name: "unsorted and overlap",
			args: args{
				request: []Booking{
					{
						Checkin: "2018-01-01",
						Nights:  5,
					},
					{
						Checkin: "2018-01-20",
						Nights:  5,
					},
					{
						Checkin: "2018-01-03",
						Nights:  5,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidBooking(tt.args.request)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_booking_getter(t *testing.T) {
	id := "A"
	checkin := "2020-05-05"
	nights := 6
	sellingRate := float64(400)
	margin := float64(13)

	gotBooking := GetBooking(id, checkin, nights, sellingRate, margin)

	expectedBooking := Booking{
		Id:          id,
		Checkin:     checkin,
		Nights:      nights,
		SellingRate: sellingRate,
		Margin:      margin,
	}

	assert.Equal(t, expectedBooking, gotBooking)
}

func TestCheckValidCombinations(t *testing.T) {
	ch := make(chan []Booking)

	combination1 := []Booking{
		{
			Id:          "A",
			Checkin:     "2020-05-01",
			Nights:      4,
			SellingRate: 600,
			Margin:      21,
		},
		{
			Id:          "B",
			Checkin:     "2020-05-09",
			Nights:      10,
			SellingRate: 530,
			Margin:      40,
		},
	}

	combination2 := []Booking{
		{
			Id:          "A",
			Checkin:     "2020-05-01",
			Nights:      8,
			SellingRate: 530,
			Margin:      40,
		},
	}

	combination3 := []Booking{
		{
			Id:          "A",
			Checkin:     "2020-05-01",
			Nights:      10,
			SellingRate: 600,
			Margin:      21,
		},
		{
			Id:          "B",
			Checkin:     "2020-05-09",
			Nights:      10,
			SellingRate: 530,
			Margin:      40,
		},
	}

	go func(ch chan []Booking) {
		defer close(ch)
		ch <- combination1
		ch <- combination2
		ch <- combination3
	}(ch)

	gotCombinations := CheckValidCombinations(ch)

	assert.Contains(t, gotCombinations, combination1)
	assert.Contains(t, gotCombinations, combination2)
	assert.NotContains(t, gotCombinations, combination3)

}
