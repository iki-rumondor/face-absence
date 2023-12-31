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

type SchoolYearHandler struct {
	Service *application.SchoolYearService
}

func NewSchoolYearHandler(service *application.SchoolYearService) *SchoolYearHandler {
	return &SchoolYearHandler{
		Service: service,
	}
}

func (h *SchoolYearHandler) GetSchoolYearPagination(c *gin.Context) {

	urlPath := c.Request.URL.Path

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "0"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))

	pagination := domain.Pagination{
		Limit: limit,
		Page:  page,
	}

	result, err := h.Service.SchoolYearPagination(urlPath, &pagination)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Berhasil mendapatkan data tahun pelajaran",
		Data:    result,
	})

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
		Message: "Berhasil menambah data tahun pelajaran",
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
		Message: "Berhasil mendapatkan data tahun pelajaran",
		Data:    res,
	})

}

func (h *SchoolYearHandler) GetSchoolYear(c *gin.Context) {

	uuid := c.Param("uuid")
	result, err := h.Service.GetSchoolYear(uuid)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	res := response.SchoolYearResponse{
		Uuid:      result.Uuid,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Berhasil mendapatkan data tahun pelajaran",
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
		Uuid: res.Uuid,
		Name: body.Name,
	}

	if err := h.Service.UpdateSchoolYear(&model); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Berhasil memperbarui data tahun pelajaran",
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
		Uuid: res.Uuid,
	}

	if err := h.Service.DeleteSchoolYear(&model); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "Berhasil menghapus data tahun pelajaran",
	})

}
