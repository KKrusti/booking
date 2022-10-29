package valueobjects

import (
	"github.com/KKrusti/booking/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CalcAverage(t *testing.T) {
	requests := Bookings{
		[]domain.Booking{
			domain.CreateBooking("bookata_XY123", "2020-01-01", 5, 200, 20),
			domain.CreateBooking("kayete_PP234", "2020-01-04", 4, 156, 22),
		},
	}

	statsCalculator := requests.CalcStats()

	expectedAverage := 8.29
	assert.Equal(t, expectedAverage, statsCalculator.AverageNight)

	expectedMinimum := 8.0
	assert.Equal(t, expectedMinimum, statsCalculator.MinimumNight)

	expectedMaximum := 8.58
	assert.Equal(t, expectedMaximum, statsCalculator.MaximumNight)
}

func Test_validBooking(t *testing.T) {
	type args struct {
		request Bookings
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "large distance",
			args: args{
				request: Bookings{Bookings: []domain.Booking{
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
			},
			want: true,
		},
		{
			name: "finish and start same day",
			args: args{
				request: Bookings{Bookings: []domain.Booking{
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
			},
			want: true,
		},
		{
			name: "overlap",
			args: args{
				request: Bookings{Bookings: []domain.Booking{
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
			},
			want: false,
		},
		{
			name: "unsorted but valid",
			args: args{
				request: Bookings{Bookings: []domain.Booking{
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
			},
			want: true,
		},
		{
			name: "unsorted and overlap",
			args: args{
				request: Bookings{Bookings: []domain.Booking{
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
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.request.isValidBooking()
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_sortByCheckinDate(t *testing.T) {
	bookings := Bookings{
		[]domain.Booking{
			domain.CreateBooking("A", "2020-01-20", 0, 0.0, 0.0),
			domain.CreateBooking("B", "2020-01-18", 0, 0.0, 0.0),
		},
	}
	bookings.sortByCheckinDate()

	assert.Equal(t, "B", bookings.Bookings[0].Id)
	assert.Equal(t, "A", bookings.Bookings[1].Id)
}

func Test_calcAverageNight(t *testing.T) {
	profits := []float64{2.2, 4.6, 8, 15.2}

	average := calcAverageNight(profits)
	expectedAverage := 7.5
	assert.Equal(t, expectedAverage, average)
}

func Test_allCombinations(t *testing.T) {

	request := Bookings{Bookings: []domain.Booking{{Id: "A"}, {Id: "B"}, {Id: "C"}}}

	ch := make(chan Bookings)
	go request.generateAllCombinations(ch)

	var combination []Bookings
	for received := range ch {
		combination = append(combination, received)
	}

	assert.Equal(t, 7, len(combination))

	//a := []Booking{{id: "A"}}
	a := Bookings{Bookings: []domain.Booking{{Id: "A"}}}
	b := Bookings{Bookings: []domain.Booking{{Id: "B"}}}
	c := Bookings{Bookings: []domain.Booking{{Id: "C"}}}
	ab := Bookings{Bookings: []domain.Booking{{Id: "A"}, {Id: "B"}}}
	bc := Bookings{Bookings: []domain.Booking{{Id: "B"}, {Id: "C"}}}
	ac := Bookings{Bookings: []domain.Booking{{Id: "A"}, {Id: "C"}}}
	abc := Bookings{Bookings: []domain.Booking{{Id: "A"}, {Id: "B"}, {Id: "C"}}}
	assert.Contains(t, combination, a)
	assert.Contains(t, combination, b)
	assert.Contains(t, combination, c)
	assert.Contains(t, combination, ab)
	assert.Contains(t, combination, ac)
	assert.Contains(t, combination, bc)
	assert.Contains(t, combination, abc)
}

func Test_CheckValidCombinations(t *testing.T) {
	ch := make(chan Bookings)

	combination1 := Bookings{Bookings: []domain.Booking{
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
	},
	}

	combination2 := Bookings{Bookings: []domain.Booking{
		{
			Id:          "A",
			Checkin:     "2020-05-01",
			Nights:      8,
			SellingRate: 530,
			Margin:      40,
		},
	},
	}

	combination3 := Bookings{Bookings: []domain.Booking{
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
	},
	}

	go func(ch chan Bookings) {
		defer close(ch)
		ch <- combination1
		ch <- combination2
		ch <- combination3
	}(ch)

	gotCombinations := filterValidCombinations(ch)

	assert.Contains(t, gotCombinations, combination1)
	assert.Contains(t, gotCombinations, combination2)
	assert.NotContains(t, gotCombinations, combination3)
}

func Test_calculateProfits(t *testing.T) {
	combinations := []Bookings{
		{
			Bookings: []domain.Booking{
				{
					Id:          "A",
					Checkin:     "2022-01-01",
					Nights:      5,
					SellingRate: 300,
					Margin:      31,
				},
				{
					Id:          "B",
					Checkin:     "2022-01-09",
					Nights:      6,
					SellingRate: 400,
					Margin:      20,
				},
			},
		},
	}

	gotStats := calculateProfits(combinations)

	expectedStats := []Stats{
		{
			RequestIds:   []string{"A", "B"},
			TotalProfit:  173,
			AverageNight: 15.97,
			MinimumNight: 13.33,
			MaximumNight: 18.6,
		},
	}

	assert.Equal(t, gotStats, expectedStats)
}

func Test_ProcessAllCombinations(t *testing.T) {
	bookings := Bookings{
		Bookings: []domain.Booking{
			{
				Id:          "A",
				Checkin:     "2020-01-01",
				Nights:      5,
				SellingRate: 200,
				Margin:      20,
			},
			{
				Id:          "B",
				Checkin:     "2020-01-04",
				Nights:      4,
				SellingRate: 156,
				Margin:      5,
			},
			{
				Id:          "C",
				Checkin:     "2020-01-04",
				Nights:      4,
				SellingRate: 150,
				Margin:      6,
			},
			{
				Id:          "D",
				Checkin:     "2020-01-10",
				Nights:      4,
				SellingRate: 160,
				Margin:      30,
			},
		},
	}

	result := bookings.ProcessAllCombinations()

	expectedStats := Stats{
		RequestIds:   []string{"A", "D"},
		TotalProfit:  88,
		AverageNight: 10,
		MinimumNight: 8,
		MaximumNight: 12,
	}

	assert.Equal(t, expectedStats, result)

}
