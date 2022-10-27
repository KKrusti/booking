package entities

import (
	utils "github.com/KKrusti/booking/internal/core"
	"github.com/stretchr/testify/assert"
	"sync"
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
	wg := &sync.WaitGroup{}
	go GenerateAllCombinations(ch, wg, request)

	var combination [][]Booking
	for received := range ch {
		wg.Done()
		combination = append(combination, received)
	}
	wg.Wait()

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
