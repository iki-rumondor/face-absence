package utils

import (
	"mime/multipart"
	"strings"

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
