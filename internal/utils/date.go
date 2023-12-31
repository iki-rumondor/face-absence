package utils

import (
	"time"
)

func IsValidDateFormat(input string) bool {
	dateFormat := "2006-01-02"
	_, err := time.Parse(dateFormat, input)
	return err == nil
}

func IsValidTimeFormat(input string) bool {
	timeFormat := "15:04"
	_, err := time.Parse(timeFormat, input)
	return err == nil
}

func IsTodayEqualTo(targetDate string) bool {
	today := time.Now().Format("2006-01-02")
	return today == targetDate
}

func IsBeforeTime(targetTime string) bool {
	parsedTime, err := time.Parse("15:04", targetTime)
	if err != nil {
		return false
	}

	now := time.Now()

	timeFormat := time.Date(now.Year(), now.Month(), now.Day(), parsedTime.Hour(), parsedTime.Minute(), 0, 0, now.Location())

	return now.Before(timeFormat)
}

func IsAfterTime(targetTime string) bool {
	parsedTime, err := time.Parse("15:04", targetTime)
	if err != nil {
		return false
	}

	now := time.Now()

	timeFormat := time.Date(now.Year(), now.Month(), now.Day(), parsedTime.Hour(), parsedTime.Minute(), 0, 0, now.Location())

	return now.After(timeFormat)
}
