package utils

import (
	"mime/multipart"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
)

func IsExcelFile(file *multipart.FileHeader) error {

	if fileExt := strings.ToLower(file.Filename[strings.LastIndex(file.Filename, ".")+1:]); fileExt != "xlsx" {
		return &response.Error{
			Code:    404,
			Message: "File uploaded is not an excel file, please use .xlsx file",
		}
	}

	return nil
}

func IsValidImageExtension(filename string) bool {
	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}

	extension := GetFileExtension(filename)
	return allowedExtensions[extension]
}

func IsValidImageSize(size int64) bool {
	const maxFileSize = 5 * 1024 * 1024
	return size <= maxFileSize
}

func InitCustomValidation() {

	govalidator.TagMap["date"] = govalidator.Validator(func(str string) bool {
		var dateFormat = "2006-01-02"
		return IsValidDate(dateFormat, str)
	})
}
