package customHTTP

import (
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/application"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
)

type ScheduleHandler struct {
	Service *application.ScheduleService
}

func NewScheduleHandler(service *application.ScheduleService) *ScheduleHandler {
	return &ScheduleHandler{
		Service: service,
	}
}

func (h *ScheduleHandler) CreateSchedule(c *gin.Context) {
	var body request.CreateSchedule
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Message: err.Error(),
		})
		return
	}

	dayFormat := "02-01-2006"
	day, err := time.Parse(dayFormat, body.Day)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Message: "Failed to parse day: " + err.Error(),
		})
		return
	}

	timeFormat := "15:04:05"
	start, err := time.Parse(timeFormat, body.Start)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Message: "Failed to parse start: " + err.Error(),
		})
		return
	}
	end, err := time.Parse(timeFormat, body.End)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Message: "Failed to parse end: " + err.Error(),
		})
		return
	}

	model := domain.Schedule{
		Uuid:         uuid.NewString(),
		Name:         body.Name,
		Day:          day,
		Start:        start,
		End:          end,
		ClassID:      body.ClassID,
		SubjectID:    body.SubjectID,
		TeacherID:    body.TeacherID,
		SchoolYearID: body.SchoolYearID,
	}

	if err := h.Service.CreateSchedule(&model); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "Schedule has been created successfully",
	})

}

func (h *ScheduleHandler) GetAllSchedules(c *gin.Context) {

	res, err := h.Service.GetAllSchedules()

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Schedule was found",
		Data:    res,
	})

}

func (h *ScheduleHandler) GetSchedule(c *gin.Context) {

	uuid := c.Param("uuid")
	res, err := h.Service.GetSchedule(uuid)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Schedule was found",
		Data:    res,
	})

}

func (h *ScheduleHandler) UpdateSchedule(c *gin.Context) {

	var body request.UpdateSchedule
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Message: err.Error(),
		})
		return
	}

	uuid := c.Param("uuid")
	res, err := h.Service.GetSchedule(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	dayFormat := "02-01-2006"
	day, err := time.Parse(dayFormat, body.Day)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Message: "Failed to parse day: " + err.Error(),
		})
		return
	}

	timeFormat := "15:04:05"
	start, err := time.Parse(timeFormat, body.Start)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Message: "Failed to parse start: " + err.Error(),
		})
		return
	}
	end, err := time.Parse(timeFormat, body.End)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Message: "Failed to parse end: " + err.Error(),
		})
		return
	}

	model := domain.Schedule{
		ID:   res.ID,
		Name:         body.Name,
		Day:          day,
		Start:        start,
		End:          end,
		ClassID:      body.ClassID,
		SubjectID:    body.SubjectID,
		TeacherID:    body.TeacherID,
		SchoolYearID: body.SchoolYearID,
	}

	if err := h.Service.UpdateSchedule(&model); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "Schedule has been updated successfully",
	})

}

func (h *ScheduleHandler) DeleteSchedule(c *gin.Context) {

	uuid := c.Param("uuid")
	res, err := h.Service.GetSchedule(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	model := domain.Schedule{
		ID: res.ID,
	}

	if err := h.Service.DeleteSchedule(&model); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "Schedule has been deleted successfully",
	})

}
