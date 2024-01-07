package customHTTP

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/application"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
)

type AbsenceHandler struct {
	Service         *application.AbsenceService
	ScheduleService *application.ScheduleService
}

func NewAbsenceHandler(service *application.AbsenceService, schedule *application.ScheduleService) *AbsenceHandler {
	return &AbsenceHandler{
		Service:         service,
		ScheduleService: schedule,
	}
}

func (h *AbsenceHandler) CreateAbsence(c *gin.Context) {

	id := c.GetUint("user_id")
	if id == 0 {
		utils.HandleError(c, &response.Error{
			Code:    500,
			Message: "Gagal menemukan id user",
		})
		return
	}

	schedule, err := h.ScheduleService.GetSchedule(c.PostForm("schedule_uuid"))
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	status, err := h.Service.CheckSchedule(schedule.ID)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	absence := domain.Absence{
		Uuid:       uuid.NewString(),
		StudentID:  id,
		ScheduleID: schedule.ID,
		Status:     status,
	}

	file, err := c.FormFile("face_image")
	if err != nil {
		utils.HandleError(c, &response.Error{
			Code:    400,
			Message: "Field face_image tidak ditemukan",
		})
	}

	if ok := utils.IsValidImageExtension(file.Filename); !ok {
		utils.HandleError(c, &response.Error{
			Code:    400,
			Message: "File yang diupload bukan sebuah gambar",
		})
		return
	}

	if ok := utils.IsValidImageSize(file.Size); !ok {
		utils.HandleError(c, &response.Error{
			Code:    400,
			Message: "File maksimal 5MB",
		})
		return
	}

	// Buat Save File Di Folder
	folder := "internal/assets/temp"
	filename := utils.GenerateRandomFileName(file.Filename)
	pathFile := filepath.Join(folder, filename)

	if err := c.SaveUploadedFile(file, pathFile); err != nil {
		utils.HandleError(c, &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan ketika menyimpan file",
		})
	}

	defer func() {
		if err := os.Remove(pathFile); err != nil {
			fmt.Println(err.Error())
		}
	}()

	if err := h.Service.CreateAbsence(&absence, pathFile); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "Absensi berhasil disimpan",
	})
}
