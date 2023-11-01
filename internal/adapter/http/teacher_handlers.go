package customHTTP

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/application"
)

type TeacherHandlers struct {
	Service *application.TeacherService
}

func NewTeacherHandler(service *application.TeacherService) *TeacherHandlers {
	return &TeacherHandlers{
		Service: service,
	}
}

func (h *TeacherHandlers) CreateTeacher(c *gin.Context){
	var body request.CreateTeacher
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
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

	if err := h.Service.CreateTeacher(body); err != nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.FailedResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "berhasil menambahkan admin guru baru",
	})

}
