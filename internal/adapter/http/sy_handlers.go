package customHTTP

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/application"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
)

type SchoolYearHandler struct {
	Service *application.SchoolYearService
}

func NewSchoolYearHandler(service *application.SchoolYearService) *SchoolYearHandler {
	return &SchoolYearHandler{
		Service: service,
	}
}

func (h *SchoolYearHandler) CreateSchoolYear(c *gin.Context) {
	var body request.CreateSchoolYear
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

	model := domain.SchoolYear{
		Uuid: uuid.NewString(),
		Name: body.Name,
	}

	if err := h.Service.CreateSchoolYear(&model); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "SchoolYear has been created successfully",
	})

}

func (h *SchoolYearHandler) GetAllSchoolYears(c *gin.Context) {

	res, err := h.Service.GetAllSchoolYears()

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "SchoolYear was found",
		Data:    res,
	})

}

func (h *SchoolYearHandler) GetSchoolYear(c *gin.Context) {

	uuid := c.Param("uuid")
	res, err := h.Service.GetSchoolYear(uuid)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "SchoolYear was found",
		Data:    res,
	})

}

func (h *SchoolYearHandler) UpdateSchoolYear(c *gin.Context) {

	var body request.UpdateSchoolYear
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
	res, err := h.Service.GetSchoolYear(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	model := domain.SchoolYear{
		ID:   res.ID,
		Name: body.Name,
	}

	if err := h.Service.UpdateSchoolYear(&model); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "SchoolYear has been updated successfully",
	})

}

func (h *SchoolYearHandler) DeleteSchoolYear(c *gin.Context) {

	uuid := c.Param("uuid")
	res, err := h.Service.GetSchoolYear(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	model := domain.SchoolYear{
		ID: res.ID,
	}

	if err := h.Service.DeleteSchoolYear(&model); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "SchoolYear has been deleted successfully",
	})

}
