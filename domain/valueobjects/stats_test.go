package valueobjects

import (
	utils "github.com/KKrusti/booking/shared"
	"github.com/stretchr/testify/assert"
	"testing"
)

//func Test_calcMinimum_first_element(t *testing.T) {
//	value := float64(1)
//	minimumValue := float64(50)
//
//	minimum := calcMinimum(0, value, minimumValue)
//
//	assert.Equal(t, minimumValue, minimum)
//
//}

//func Test_CalcMinimum(t *testing.T) {
//	type args struct {
//		index           int
//		value           float64
//		previousMinimum float64
//	}
//	tests := []struct {
//		name string
//		args args
//		want float64
//	}{
//		{
//			name: "first comparison",
//			args: args{
//				index:           0,
//				value:           float64(20),
//				previousMinimum: float64(0),
//			},
//			want: float64(20),
//		},
//		{
//			name: "not first value and comparison previous value smaller",
//			args: args{
//				index:           4,
//				value:           float64(20),
//				previousMinimum: float64(10),
//			},
//			want: float64(10),
//		},
//		{
//			name: "not first value and comparison previous value bigger",
//			args: args{
//				index:           4,
//				value:           float64(10),
//				previousMinimum: float64(20),
//			},
//			want: float64(10),
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got := calcMinimum(tt.args.index, tt.args.previousMinimum, tt.args.value)
//			assert.Equal(t, tt.want, got)
//		})
//	}
//}
//
//func Test_CalcMaximum(t *testing.T) {
//	type args struct {
//		value           float64
//		previousMaximum float64
//	}
//	tests := []struct {
//		name string
//		args args
//		want float64
//	}{
//		{
//			name: "previous value bigger",
//			args: args{
//				value:           float64(10),
//				previousMaximum: float64(20),
//			},
//			want: float64(20),
//		},
//		{
//			name: "current value bigger",
//			args: args{
//				value:           float64(20),
//				previousMaximum: float64(10),
//			},
//			want: float64(20),
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got := calcMaximum(tt.args.previousMaximum, tt.args.value)
//			assert.Equal(t, tt.want, got)
//		})
//	}
//}

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
