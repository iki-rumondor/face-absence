package utils

import (
	"fmt"
	"strings"
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

func GetHariIndonesia(englishDay string) string {
	hariMapping := map[string]string{
		"Monday":    "Senin",
		"Tuesday":   "Selasa",
		"Wednesday": "Rabu",
		"Thursday":  "Kamis",
		"Friday":    "Jumat",
		"Saturday":  "Sabtu",
		"Sunday":    "Minggu",
	}

	return hariMapping[englishDay]
}

func IsDayEqualTo(dayString string) bool {
	hariInggris := time.Now().Weekday().String()
	hariIndonesia := GetHariIndonesia(hariInggris)

	return strings.EqualFold(hariIndonesia, dayString)
}

func IsBeforeTime(targetTime string) bool {
	parsedTime, err := time.Parse("15:04", targetTime)
	if err != nil {
		return false
	}

	now := time.Now()
	witaLocation, err := time.LoadLocation("Asia/Makassar")
	if err != nil {
		return false
	}

	witaTime := now.In(witaLocation)

	timeFormat := time.Date(witaTime.Year(), witaTime.Month(), witaTime.Day(), parsedTime.Hour(), parsedTime.Minute(), 0, 0, witaTime.Location())
	fmt.Println(witaTime, timeFormat)
	return witaTime.Before(timeFormat)
}

func IsAfterTime(targetTime string) bool {
	parsedTime, err := time.Parse("15:04", targetTime)
	if err != nil {
		return false
	}

	now := time.Now()
	witaLocation, err := time.LoadLocation("Asia/Makassar")
	if err != nil {
		return false
	}
	witaTime := now.In(witaLocation)

	timeFormat := time.Date(witaTime.Year(), witaTime.Month(), witaTime.Day(), parsedTime.Hour(), parsedTime.Minute(), 0, 0, witaTime.Location())

	return witaTime.After(timeFormat)
}
