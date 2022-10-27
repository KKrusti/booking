package services

import (
	utils "github.com/KKrusti/booking/internal/core"
	"github.com/KKrusti/booking/internal/core/domain/entities"
	"sort"
	"time"
)

func Maximize(request []entities.Request) entities.MaxResponse {
	return entities.MaxResponse{}
}

func isValidCombination(requests []entities.Request) bool {
	sortByCheckinDate(requests)
	for i := 0; i < len(requests)-1; i++ {
		currentCheckout := getCheckoutDate(requests[i])
		nextCheckin := utils.StringToTime(requests[i+1].Checkin)
		if nextCheckin.Before(currentCheckout) {
			return false
		}
	}
	return true
}

func sortByCheckinDate(requests []entities.Request) {
	sort.Slice(requests[:], func(i, j int) bool {
		return requests[i].Checkin < requests[j].Checkin
	})
}

func getCheckoutDate(request entities.Request) time.Time {
	checkinDate := utils.StringToTime(request.Checkin)
	checkoutDate := checkinDate.AddDate(0, 0, request.Nights)
	return checkoutDate
}
