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

type ClassHandler struct {
	Service *application.ClassService
}

func NewClassHandler(service *application.ClassService) *ClassHandler {
	return &ClassHandler{
		Service: service,
	}
}

func (h *ClassHandler) CreateClass(c *gin.Context) {
	var body request.CreateClass
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

	class := domain.Class{
		Uuid:      uuid.NewString(),
		Name:      body.Name,
		TeacherID: body.TeacherID,
	}

	if err := h.Service.CreateClass(&class); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "New class has been saved successfully",
	})

}

func (h *ClassHandler) GetAllClasses(c *gin.Context) {

	classes, err := h.Service.GetAllClasses()

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Success to find all classes",
		Data:    classes,
	})

}

func (h *ClassHandler) GetClassPagination(c *gin.Context) {

	urlPath := c.Request.URL.Path

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))

	pagination := domain.Pagination{
		Limit: limit,
		Page:  page,
	}

	result, err := h.Service.ClassPagination(urlPath, &pagination)

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

func (h *ClassHandler) GetClass(c *gin.Context) {

	uuid := c.Param("uuid")
	class, err := h.Service.GetClass(uuid)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Success to find class by uuid",
		Data:    class,
	})

}

func (h *ClassHandler) UpdateClass(c *gin.Context) {

	var body request.UpdateClass
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
	classInDB, err := h.Service.GetClass(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	class := domain.Class{
		Uuid:      classInDB.Uuid,
		Name:      body.Name,
		TeacherID: body.TeacherID,
	}

	if err := h.Service.UpdateClass(&class); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Success to update class by uuid",
	})

}

func (h *ClassHandler) DeleteClass(c *gin.Context) {

	uuid := c.Param("uuid")
	classInDB, err := h.Service.GetClass(uuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	class := domain.Class{
		Uuid: classInDB.Uuid,
	}

	if err := h.Service.DeleteClass(&class); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Success to delete class by uuid",
	})

}
