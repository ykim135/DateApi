package service

import (
	"time"
)

func GetFirstMonday(month time.Month) time.Time {
	firstDayOfMonth := time.Date(2017, month, 1, 0, 0, 0, 0, time.UTC)

	daysLeftUntilFirstMonday := (7 - int(firstDayOfMonth.Weekday()) + 1) % 7

	return firstDayOfMonth.AddDate(0, 0, daysLeftUntilFirstMonday)
}

func GetAllFirstMonday(allFirstMonday []time.Time, month time.Month) []time.Time {
	var firstMondayOfMonth time.Time = GetFirstMonday(month)

	var newAllFirstMonday []time.Time = append(allFirstMonday, firstMondayOfMonth)

	if month == 12 {
		return newAllFirstMonday
	}
	return GetAllFirstMonday(newAllFirstMonday, month+1)
}
