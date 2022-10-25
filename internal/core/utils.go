package utils

import "math"

func Round(number float64) float64 {
	return math.Round(number*100) / 100
}
