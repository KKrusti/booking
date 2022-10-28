package shared

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
			got := Round(tt.args.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
