package services

import (
	"github.com/KKrusti/booking/internal/core/domain/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_valid_combinations(t *testing.T) {
	type args struct {
		request []entities.Booking
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "large distance",
			args: args{
				request: []entities.Booking{
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
				request: []entities.Booking{
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
				request: []entities.Booking{
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
				request: []entities.Booking{
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
				request: []entities.Booking{
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
			got := entities.IsValidCombination(tt.args.request)
			assert.Equal(t, tt.want, got)
		})
	}
}
