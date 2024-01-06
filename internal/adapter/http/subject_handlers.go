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

type SubjectHandler struct {
	Service *application.SubjectService
}

func NewSubjectHandler(service *application.SubjectService) *SubjectHandler {
	return &SubjectHandler{
		Service: service,
	}
}

func (h *SubjectHandler) CreateSubject(c *gin.Context) {
	var body request.CreateSubject
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

	model := domain.Subject{
		Uuid: uuid.NewString(),
		Name: body.Name,
	}

	if err := h.Service.CreateSubject(&model); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "Subject has been created successfully",
	})

}

func (h *SubjectHandler) GetSubjectPagination(c *gin.Context) {

	urlPath := c.Request.URL.Path

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "0"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))

	pagination := domain.Pagination{
		Limit: limit,
		Page:  page,
	}

	result, err := h.Service.SubjectPagination(urlPath, &pagination)

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

func (h *SubjectHandler) GetAllSubjects(c *gin.Context) {

	res, err := h.Service.GetAllSubjects()

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Subject was found",
		Data:    res,
	})

}

func (h *SubjectHandler) GetSubject(c *gin.Context) {

	uuid := c.Param("uuid")
	res, err := h.Service.GetSubject(uuid)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Subject was found",
		Data:    res,
	})

}

func (h *SubjectHandler) UpdateSubject(c *gin.Context) {

	var body request.UpdateSubject
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
	res, err := h.Service.GetSubject(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	model := domain.Subject{
		Uuid: res.Uuid,
		Name: body.Name,
	}

	if err := h.Service.UpdateSubject(&model); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "Subject has been updated successfully",
	})

}

func (h *SubjectHandler) DeleteSubject(c *gin.Context) {

	uuid := c.Param("uuid")
	res, err := h.Service.GetSubject(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	model := domain.Subject{
		Uuid: res.Uuid,
	}

	if err := h.Service.DeleteSubject(&model); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "Subject has been deleted successfully",
	})

}
