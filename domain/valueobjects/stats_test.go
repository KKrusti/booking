package valueobjects

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsMoreProfitable(t *testing.T) {
	type args struct {
		currentStats Stats
		newStats     Stats
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "current is more profitable",
			args: args{
				currentStats: Stats{TotalProfit: float64(30)},
				newStats:     Stats{TotalProfit: float64(10)},
			},
			want: true,
		},
		{
			name: "newer is more profitable",
			args: args{
				currentStats: Stats{TotalProfit: float64(10)},
				newStats:     Stats{TotalProfit: float64(30)},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.currentStats.isMoreProfitableThan(tt.args.newStats)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_sortByMostProfitable(t *testing.T) {
	lessProfitIds := []string{"A", "B"}
	moreProfitIds := []string{"C", "D"}

	calcsToSort := []Stats{
		{
			RequestIds:  lessProfitIds,
			TotalProfit: 40,
		},
		{
			RequestIds:  moreProfitIds,
			TotalProfit: 53,
		},
	}

	//before sorting
	assert.Equal(t, calcsToSort[0].RequestIds, lessProfitIds)
	assert.Equal(t, calcsToSort[1].RequestIds, moreProfitIds)

	sortByMostProfitableBooking(calcsToSort)

	assert.Equal(t, calcsToSort[0].RequestIds, moreProfitIds)
	assert.Equal(t, calcsToSort[1].RequestIds, lessProfitIds)
}
