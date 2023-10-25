package customHTTP

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/application"
)

type StudentHandlers struct {
	Service *application.StudentService
}

func NewStudentHandler(service *application.StudentService) *StudentHandlers {
	return &StudentHandlers{
		Service: service,
	}
}

func (h *StudentHandlers) ImportStudentsData(c *gin.Context) {
	file, _ := c.FormFile("students")
	tempFolder := "internal/temp"
	pathFile := filepath.Join(tempFolder, file.Filename)

	if err := c.SaveUploadedFile(file, pathFile); err != nil {
		fmt.Println(err.Error())
	}

	defer func() {
		if err := os.Remove(pathFile); err != nil {
			fmt.Println(err.Error())
		}
	}()

	failed, err := h.Service.ImportStudents(pathFile)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.FailedResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse{
		Success: true,
		Message: "students has been saved successfully",
		Data:    failed,
	})
}

func (h *StudentHandlers) GetAllStudentsData(c *gin.Context) {
	users, err := h.Service.GetAllStudentUsers()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.FailedResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "get all students has been successfully",
		Data:    users,
	})
}

func (h *StudentHandlers) GetStudentData(c *gin.Context) {

	student, err := h.Service.GetStudentUser(c.Param("uuid"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.FailedResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: fmt.Sprintf("student with uuid %s is found", student.Uuid),
		Data:    student,
	})
}

func (h *StudentHandlers) UpdateStudentData(c *gin.Context) {

	body, ok := c.Get("body")
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.FailedResponse{
			Success: false,
			Message: "something went wrong at internal server",
		})
	}

	var student request.Student = body.(request.Student)

	result, err := h.Service.GetStudent(c.Param("uuid"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.FailedResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	student.ID = result.ID
	student.UserID = result.UserID

	if err := h.Service.UpdateStudentUser(&student); err != nil {
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: fmt.Sprintf("student with nis %s has been updated successfully", student.NIS),
	})
}

func (h *StudentHandlers) DeleteStudentData(c *gin.Context) {

	result, err := h.Service.GetStudent(c.Param("uuid"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.FailedResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	if err := h.Service.DeleteStudentUser(result); err != nil {
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: fmt.Sprintf("student with nis %s has been updated successfully", result.NIS),
	})
}
