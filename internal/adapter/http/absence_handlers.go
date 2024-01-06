package customHTTP

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/application"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
)

type AbsenceHandler struct {
	Service *application.AbsenceService
}

func NewAbsenceHandler(service *application.AbsenceService) *AbsenceHandler {
	return &AbsenceHandler{
		Service: service,
	}
}

func (h *AbsenceHandler) CreateAbsence(c *gin.Context) {

	id := c.GetUint("user_id")
	if id == 0 {
		utils.HandleError(c, &response.Error{
			Code:    500,
			Message: "Can't get user id",
		})
		return
	}

	scheduleID, err := strconv.Atoi(c.PostForm("schedule_id"))
	if err != nil{
		utils.HandleError(c, &response.Error{
			Code:    400,
			Message: "Schedule id is not a number",
		})
		return
	}

	status, err := h.Service.CheckSchedule(uint(scheduleID))
	
	if err != nil{
		utils.HandleError(c, err)
		return
	}

	absence := domain.Absence{
		Uuid: uuid.NewString(),
		StudentID: id,
		ScheduleID: uint(scheduleID),
		Status: status,
	}

	file, err := c.FormFile("face_image")
	if err != nil {
		utils.HandleError(c, &response.Error{
			Code:    400,
			Message: "Face image is not found",
		})
	}


	if ok := utils.IsValidImageExtension(file.Filename); !ok {
		utils.HandleError(c, &response.Error{
			Code:    400,
			Message: "File uploaded is not an image",
		})
		return
	}

	if ok := utils.IsValidImageSize(file.Size); !ok {
		utils.HandleError(c, &response.Error{
			Code:    400,
			Message: "File size limit: 5MB",
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
			Message: "Something went wrong when uploaded file",
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
		Message: "absence has been created successfully",
	})
}
