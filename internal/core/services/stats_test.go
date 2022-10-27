package services

import (
	utils "github.com/KKrusti/booking/internal/core"
	"github.com/KKrusti/booking/internal/core/domain/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CalcAverage(t *testing.T) {
	request1 := entities.GetBooking("bookata_XY123", "2020-01-01", 5, 200, 20)

	request2 := entities.GetBooking("kayete_PP234", "2020-01-04", 4, 156, 22)
	requests := []entities.Booking{request1, request2}

	statsCalculator := CalcStats(requests)

	expectedAverage := 8.29
	assert.Equal(t, expectedAverage, statsCalculator.AverageNight)

	expectedMinimum := 8.0
	assert.Equal(t, expectedMinimum, statsCalculator.MinimumNight)

	expectedMaximum := 8.58
	assert.Equal(t, expectedMaximum, statsCalculator.MaximumNight)
}

func Test_calcAverageNight(t *testing.T) {
	profits := []float64{2.2, 4.6, 8, 15.2}

	average := calcAverageNight(profits)
	expectedAverage := 7.5
	assert.Equal(t, expectedAverage, average)
}

func Test_calcMinimum_first_element(t *testing.T) {
	value := float64(1)
	minimumValue := float64(50)

	minimum := calcMinimum(0, value, minimumValue)

	assert.Equal(t, minimumValue, minimum)

}

func Test_calcMinimum(t *testing.T) {
	value := float64(1)
	previousMinimum := float64(50)

	minimum := calcMinimum(1, value, previousMinimum)

	assert.Equal(t, value, minimum)
}

func Test_CalcMaximum(t *testing.T) {
	value := float64(20)
	previousMaximum := float64(10)

	maximum := calcMaximum(previousMaximum, value)

	assert.Equal(t, value, maximum)
}

func Test_round(t *testing.T) {
	type args struct {
		n float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "lower",
			args: args{
				n: 8.5122222,
			},
			want: 8.51,
		},
		{
			name: "upper",
			args: args{
				n: 8.5166666,
			},
			want: 8.52,
		},
		{
			name: "upper from mid",
			args: args{
				n: 8.515,
			},
			want: 8.52,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.Round(tt.args.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
