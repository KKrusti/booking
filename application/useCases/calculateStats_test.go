package useCases

import (
	"github.com/KKrusti/booking/domain"
	"github.com/KKrusti/booking/domain/valueobjects"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CalculateStats(t *testing.T) {
	type args struct {
		bookings valueobjects.Bookings
	}
	tests := []struct {
		name string
		args args
		want valueobjects.Stats
	}{
		{
			name: "sample 1",
			args: args{
				bookings: valueobjects.Bookings{
					[]domain.Booking{
						{
							Id:          "A",
							Checkin:     "2020-01-01",
							Nights:      5,
							SellingRate: 200,
							Margin:      15,
						},
						{
							Id:          "B",
							Checkin:     "2020-01-05",
							Nights:      4,
							SellingRate: 150,
							Margin:      33,
						},
					},
				},
			},
			want: valueobjects.Stats{
				AverageNight: 9.19,
				MinimumNight: 6,
				MaximumNight: 12.38,
			},
		},
		{
			name: "sample 2",
			args: args{
				bookings: valueobjects.Bookings{
					[]domain.Booking{
						{
							Id:          "A",
							Checkin:     "2020-01-01",
							Nights:      5,
							SellingRate: 200,
							Margin:      15,
						},
						{
							Id:          "B",
							Checkin:     "2020-01-05",
							Nights:      4,
							SellingRate: 150,
							Margin:      33,
						},
						{
							Id:          "C",
							Checkin:     "2020-01-010",
							Nights:      10,
							SellingRate: 3000,
							Margin:      40,
						},
					},
				},
			},
			want: valueobjects.Stats{
				AverageNight: 46.13,
				MinimumNight: 6,
				MaximumNight: 120,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateStats(tt.args.bookings)
			assert.Equal(t, tt.want.MinimumNight, got.MinimumNight)
			assert.Equal(t, tt.want.MaximumNight, got.MaximumNight)
			assert.Equal(t, tt.want.AverageNight, got.AverageNight)
		})
	}
}
