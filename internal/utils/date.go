package utils

import "time"

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
