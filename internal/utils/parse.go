package utils

import "time"

func ParseTimeTo(timeString, format string) (string, error) {
	parsedTime, err := time.Parse(time.RFC3339Nano, timeString)
	if err != nil {
		return "", err
	}

	formattedTime := parsedTime.Format(format)
	return formattedTime, nil
}
