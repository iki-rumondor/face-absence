package utils

import (
	"fmt"
	"strings"

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
	newName := fmt.Sprintf("%s.%s", u, ext)
	return newName
}
