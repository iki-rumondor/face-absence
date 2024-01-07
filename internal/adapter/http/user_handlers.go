package customHTTP

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/application"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
)

type UserHandler struct {
	Service *application.UserService
}

func NewUserHandler(service *application.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (h *UserHandler) UpdateAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		utils.HandleError(c, &response.Error{
			Code:    400,
			Message: "Field avatar tidak ditemukan",
		})
		return
	}

	id := c.GetUint("user_id")
	if id == 0 {
		utils.HandleError(c, &response.Error{
			Code:    500,
			Message: "Gagal mendapatkan id user",
		})
		return
	}

	if ok := utils.IsValidImageExtension(file.Filename); !ok {
		utils.HandleError(c, &response.Error{
			Code:    400,
			Message: "File yang diupload bukan gambar",
		})
		return
	}

	if ok := utils.IsValidImageSize(file.Size); !ok {
		utils.HandleError(c, &response.Error{
			Code:    400,
			Message: "Ukuran file maksimal 5MB",
		})
		return
	}

	// Buat Save File Di Folder
	folder := "internal/assets/avatar"
	filename := utils.GenerateRandomFileName(file.Filename)
	pathFile := filepath.Join(folder, filename)

	if err := c.SaveUploadedFile(file, pathFile); err != nil {
		utils.HandleError(c, &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		})
	}

	model := domain.User{
		ID:     id,
		Avatar: &filename,
	}

	if err := h.Service.UpdateAvatar(&model); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Berhasil mengubah avatar user",
	})
}
