package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

func GetFileExtension(filename string) string {
	lastDotIndex := strings.LastIndex(filename, ".")
	if lastDotIndex == -1 {
		return ""
	}

	return filename[lastDotIndex:]
}

func GenerateRandomFileName(filename string) string {
	u := uuid.NewString()
	ext := GetFileExtension(filename)
	newName := fmt.Sprintf("%s%s", u, ext)
	return newName
}

func GenerateStartEndDay(dayString string) (string, string) {
	day, err := time.Parse("2006-01-02", dayString)
	if err != nil {
		return "", ""
	}

	startOfDay := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, time.UTC)
	endOfDay := startOfDay.Add(24 * time.Hour).Add(-time.Nanosecond)

	formattedStartDate := startOfDay.Format("2006-01-02 15:04:05.999")
	formattedEndDate := endOfDay.Format("2006-01-02 15:04:05.999")

	return formattedStartDate, formattedEndDate
}
