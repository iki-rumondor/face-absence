package customHTTP

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/application"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
)

type StudentHandlers struct {
	Service *application.StudentService
}

func NewStudentHandler(service *application.StudentService) *StudentHandlers {
	return &StudentHandlers{
		Service: service,
	}
}

func (h *StudentHandlers) CreateStudent(c *gin.Context) {
	var body request.CreateStudent
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

	if err := h.Service.CreateStudent(&body); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "Teacher has been saved successfully",
	})

}

func (h *StudentHandlers) ImportStudentsData(c *gin.Context) {
	file, err := c.FormFile("students")
	if err != nil {
		utils.HandleError(c, &response.Error{
			Code:    400,
			Message: "Your request file is not valid",
		})
		return
	}

	if err := utils.IsExcelFile(file); err != nil {
		utils.HandleError(c, err)
		return
	}

	tempFolder := "internal/temp"
	pathFile := filepath.Join(tempFolder, file.Filename)

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

	failed, err := h.Service.ImportStudents(pathFile)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "students has been saved successfully",
		Data:    failed,
	})
}

func (h *StudentHandlers) GetAllStudentsData(c *gin.Context) {

	urlPath := c.Request.URL.Path

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))

	pagination := domain.Pagination{
		Limit: limit,
		Page:  page,
	}

	result, err := h.Service.StudentsPagination(urlPath, &pagination)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "get all students has been successfully",
		Data:    result,
	})
}

func (h *StudentHandlers) GetStudentData(c *gin.Context) {

	student, err := h.Service.GetStudent(c.Param("uuid"))

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: fmt.Sprintf("student with uuid %s is found", student.Uuid),
		Data:    student,
	})
}

func (h *StudentHandlers) UpdateStudentData(c *gin.Context) {

	var body request.UpdateStudent
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

	uuid := c.Param("uuid")

	student := domain.Student{
		Uuid:         uuid,
		NIS:          body.NIS,
		JK:           body.JK,
		TempatLahir:  body.TempatLahir,
		TanggalLahir: body.TanggalLahir,
		Alamat:       body.Alamat,
		ClassID:      body.ClassID,
	}

	user := domain.User{
		Nama:     body.Nama,
		Username: body.Username,
	}

	if err := h.Service.UpdateStudent(&student, &user); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: fmt.Sprintf("student with uuid %s has been updated successfully", uuid),
	})
}

func (h *StudentHandlers) DeleteStudent(c *gin.Context) {

	uuid := c.Param("uuid")

	if err := h.Service.DeleteStudent(uuid); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: fmt.Sprintf("student with uuid %s has been deleted successfully", uuid),
	})
}
