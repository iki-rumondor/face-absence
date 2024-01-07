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
	Service        *application.ClassService
	TeacherService *application.TeacherService
}

func NewClassHandler(service *application.ClassService, teacher *application.TeacherService) *ClassHandler {
	return &ClassHandler{
		Service:        service,
		TeacherService: teacher,
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

	teacher, err := h.TeacherService.GetTeacher(body.TeacherUuid)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	class := domain.Class{
		Uuid:      uuid.NewString(),
		Name:      body.Name,
		TeacherID: teacher.ID,
	}

	if err := h.Service.CreateClass(&class); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "Kelas baru berhasil ditambahkan",
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
		Message: "Berhasil mendapatkan seluruh kelas",
		Data:    classes,
	})

}

func (h *ClassHandler) GetClassOption(c *gin.Context) {

	classes, err := h.Service.GetClassOptions()

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Berhasil mendapatkan seluruh kelas",
		Data:    classes,
	})

}

func (h *ClassHandler) GetClassPagination(c *gin.Context) {

	urlPath := c.Request.URL.Path

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "0"))
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
		Message: "Berhasil mendapatkan kelas",
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
		Message: "Berhasil mendapatkan kelas",
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

	teacher, err := h.TeacherService.GetTeacher(body.TeacherUuid)
	if err != nil {
		utils.HandleError(c, err)
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
		TeacherID: teacher.ID,
	}

	if err := h.Service.UpdateClass(&class); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Berhasil mengupdate data kelas",
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
		Message: "Berhasil menghapus data kelas",
	})

}
