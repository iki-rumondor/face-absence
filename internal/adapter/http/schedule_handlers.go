package customHTTP

import (
	"net/http"
	"strconv"

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

func (h *ScheduleHandler) GetSchedulePagination(c *gin.Context) {

	urlPath := c.Request.URL.Path

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "0"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))

	pagination := domain.Pagination{
		Limit: limit,
		Page:  page,
	}

	result, err := h.Service.SchedulePagination(urlPath, &pagination)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "your request has been executed successfully",
		Data:    result,
	})

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

	if ok := utils.IsValidDateFormat(body.Day); !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Success: false,
			Message: "Failed to parse day, please use format yyyy-mm-dd",
		})
		return
	}

	if ok := utils.IsValidTimeFormat(body.Start); !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Success: false,
			Message: "Failed to parse start, please use format hh:mm",
		})
		return
	}

	if ok := utils.IsValidTimeFormat(body.End); !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Success: false,
			Message: "Failed to parse end, please use format hh:mm",
		})
		return
	}

	model := domain.Schedule{
		Uuid:         uuid.NewString(),
		Name:         body.Name,
		Day:          body.Day,
		Start:        body.Start,
		End:          body.End,
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

	if ok := utils.IsValidDateFormat(body.Day); !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Success: false,
			Message: "Failed to parse day, please use format yyyy-mm-dd",
		})
		return
	}

	if ok := utils.IsValidTimeFormat(body.Start); !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Success: false,
			Message: "Failed to parse start, please use format hh:mm",
		})
		return
	}

	if ok := utils.IsValidTimeFormat(body.End); !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Success: false,
			Message: "Failed to parse end, please use format hh:mm",
		})
		return
	}

	model := domain.Schedule{
		Uuid:         res.Uuid,
		Name:         body.Name,
		Day:          body.Day,
		Start:        body.Start,
		End:          body.End,
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
		Uuid: res.Uuid,
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
