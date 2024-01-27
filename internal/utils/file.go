package utils

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
)

func SaveUploadedImage(c *gin.Context) (string, error) {
	file, err := c.FormFile("image")
	if err != nil {
		return "", &response.Error{
			Code:    400,
			Message: "Field image tidak ditemukan",
		}
	}

	if ok := IsValidImageExtension(file.Filename); !ok {
		return "", &response.Error{
			Code:    400,
			Message: "File yang diupload bukan sebuah gambar",
		}
	}

	if ok := IsValidImageSize(file.Size); !ok {
		return "", &response.Error{
			Code:    400,
			Message: "File maksimal 5MB",
		}
	}

	folder := "internal/assets/avatar"
	filename := GenerateRandomFileName(file.Filename)
	pathFile := filepath.Join(folder, filename)

	if err := c.SaveUploadedFile(file, pathFile); err != nil {
		return "", &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan ketika menyimpan file",
		}
	}

	return filename, nil
}
