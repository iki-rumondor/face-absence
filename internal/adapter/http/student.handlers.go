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
	"github.com/iki-rumondor/init-golang-service/internal/domain"
)

type StudentHandlers struct {
	Service *application.StudentService
}

func NewStudentHandler(service *application.StudentService) *StudentHandlers {
	return &StudentHandlers{
		Service: service,
	}
}

func (h *StudentHandlers) CreateStudentsData(c *gin.Context) {
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

	if err := h.Service.CreateStudent(pathFile); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.StudentResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	c.JSON(http.StatusCreated, response.StudentResponse{
		Success: true,
		Message: "students has been saved successfully",
	})
}

func (h *StudentHandlers) GetAllStudentsData(c *gin.Context) {
	students, err := h.Service.GetAllStudents()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.StudentResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	var res = []*response.Student{}

	for _, student := range students.Students {
		res = append(res, &response.Student{
			ID:        student.ID,
			Uuid:      student.Uuid,
			Nama:      student.Nama,
			NIS:       student.NIS,
			Kelas:     student.Kelas,
			CreatedAt: student.CreatedAt,
			UpdatedAt: student.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response.StudentResponse{
		Success: true,
		Message: "get all students has been successfully",
		Data:    res,
	})
}

func (h *StudentHandlers) GetStudentData(c *gin.Context) {

	student := h.GetStudentDomainByUuid(c)

	res := domain.Student{
		ID:        student.ID,
		Uuid:      student.Uuid,
		Nama:      student.Nama,
		NIS:       student.NIS,
		Kelas:     student.Kelas,
		CreatedAt: student.CreatedAt,
		UpdatedAt: student.UpdatedAt,
	}

	c.JSON(http.StatusOK, response.StudentResponse{
		Success: true,
		Message: fmt.Sprintf("student with uuid %s is found", res.Uuid),
		Data:    res,
	})
}

func (h *StudentHandlers) UpdateStudentData(c *gin.Context) {

	body, ok := c.Get("body")
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.StudentResponse{
			Success: false,
			Message: "something went wrong at internal server",
		})
	}

	var req request.Student = body.(request.Student)

	student := h.GetStudentDomainByUuid(c)
	student.Nama = req.Nama
	student.NIS = req.NIS
	student.Kelas = req.Kelas

	if err := h.Service.UpdateStudent(student); err != nil {
		return
	}

	c.JSON(http.StatusOK, response.StudentResponse{
		Success: true,
		Message: fmt.Sprintf("student with uuid %s has been updated successfully", student.Uuid),
	})
}

func (h *StudentHandlers) DeleteStudentData(c *gin.Context) {

	student := h.GetStudentDomainByUuid(c)
	if err := h.Service.DeleteStudent(student); err != nil {
		return
	}

	c.JSON(http.StatusOK, response.StudentResponse{
		Success: true,
		Message: fmt.Sprintf("student with uuid %s has been deleted successfully", student.Uuid),
	})
}

func (h *StudentHandlers) GetStudentDomainByUuid(c *gin.Context) *domain.Student {
	uuid := c.Param("uuid")

	student, err := h.Service.GetStudent(uuid)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.StudentResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	return student

}
