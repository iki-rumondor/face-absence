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
	Service           *application.ScheduleService
	ClassService      *application.ClassService
	SubjectService    *application.SubjectService
	TeacherService    *application.TeacherService
	SchoolYearService *application.SchoolYearService
}

func NewScheduleHandler(service *application.ScheduleService, class *application.ClassService, subject *application.SubjectService, teacher *application.TeacherService, sy *application.SchoolYearService) *ScheduleHandler {
	return &ScheduleHandler{
		Service:           service,
		ClassService:      class,
		SubjectService:    subject,
		TeacherService:    teacher,
		SchoolYearService: sy,
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
		Message: "Berhasil menemukan jadwal",
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

	class, err := h.ClassService.GetClass(body.ClassUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	subject, err := h.SubjectService.GetSubject(body.SubjectUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	schoolYear, err := h.SchoolYearService.GetSchoolYear(body.SchoolYearUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	if ok := utils.IsValidTimeFormat(body.Start); !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Success: false,
			Message: "Gunakan format hh:mm pada field start",
		})
		return
	}

	if ok := utils.IsValidTimeFormat(body.End); !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Success: false,
			Message: "Gunakan format hh:mm pada field end",
		})
		return
	}

	model := domain.Schedule{
		Uuid:         uuid.NewString(),
		Day:          body.Day,
		Start:        body.Start,
		End:          body.End,
		ClassID:      class.ID,
		SubjectID:    subject.ID,
		SchoolYearID: schoolYear.ID,
	}

	if err := h.Service.CreateSchedule(&model); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "Jadwal berhasil ditambahkan",
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
		Message: "Jadwal berhasil ditemukan",
		Data:    res,
	})

}

func (h *ScheduleHandler) GetSchedule(c *gin.Context) {

	uuid := c.Param("uuid")
	schedule, err := h.Service.GetSchedule(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	res := response.ScheduleResponse{
		Uuid:  schedule.Uuid,
		Day:   schedule.Day,
		Start: schedule.Start,
		End:   schedule.End,
		Class: &response.ClassData{
			Uuid:      schedule.Class.Uuid,
			Name:      schedule.Class.Name,
			CreatedAt: schedule.Class.CreatedAt,
			UpdatedAt: schedule.Class.UpdatedAt,
		},
		Subject: &response.SubjectResponse{
			Uuid:      schedule.Subject.Uuid,
			Name:      schedule.Subject.Name,
			CreatedAt: schedule.Subject.CreatedAt,
			UpdatedAt: schedule.Subject.UpdatedAt,
		},
		SchoolYear: &response.SchoolYearResponse{
			Uuid:      schedule.SchoolYear.Uuid,
			Name:      schedule.SchoolYear.Name,
			CreatedAt: schedule.SchoolYear.CreatedAt,
			UpdatedAt: schedule.SchoolYear.UpdatedAt,
		},
		CreatedAt: schedule.CreatedAt,
		UpdatedAt: schedule.UpdatedAt,
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Jadwal berhasil ditemukan",
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

	class, err := h.ClassService.GetClass(body.ClassUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	subject, err := h.SubjectService.GetSubject(body.SubjectUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	schoolYear, err := h.SchoolYearService.GetSchoolYear(body.SchoolYearUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	if ok := utils.IsValidTimeFormat(body.Start); !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Success: false,
			Message: "Gunakan format hh:mm pada field start",
		})
		return
	}

	if ok := utils.IsValidTimeFormat(body.End); !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Success: false,
			Message: "Gunakan format hh:mm pada field end",
		})
		return
	}

	model := domain.Schedule{
		Uuid:         res.Uuid,
		Day:          body.Day,
		Start:        body.Start,
		End:          body.End,
		ClassID:      class.ID,
		SubjectID:    subject.ID,
		SchoolYearID: schoolYear.ID,
	}

	if err := h.Service.UpdateSchedule(&model); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Jadwal berhasil diperbarui",
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
		Message: "Jadwal berhasil dihapus",
	})

}
