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

func FormatToTime(timeString, format string) (time.Time, error) {
	parsedTime, err := time.Parse(format, timeString)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}
