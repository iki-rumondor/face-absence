package customHTTP

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
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

func (h *StudentHandlers) CreateStudent(c *gin.Context) {
	file, _ := c.FormFile("students")
	tempFolder := "internal/temp"
	pathFile := filepath.Join(tempFolder, file.Filename)

	if err := c.SaveUploadedFile(file, pathFile); err != nil {
		fmt.Println(err.Error())
	}

	defer func() {
		if err := os.Remove(pathFile); err != nil{
			fmt.Println(err.Error())
		}
	}()

	if err := h.Service.CreateStudent(pathFile); err != nil {
		fmt.Println(err.Error())
	}
}
