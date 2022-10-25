package utils

import (
	"fmt"
	"math"
	"time"
)

const layout = "2006-01-02"

func Round(number float64) float64 {
	return math.Round(number*100) / 100
}

func StringToTime(date string) time.Time {
	timeDate, err := time.Parse(layout, date)
	if err != nil {
		panic(fmt.Sprintf("cannot parse date %s", date))
	}
	return timeDate
}
