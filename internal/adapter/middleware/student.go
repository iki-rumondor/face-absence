package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
)

func ValidateStudentJSON() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body request.Student

		if err := c.BindJSON(&body); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, response.StudentResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		if _, err := govalidator.ValidateStruct(&body); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, response.StudentResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		c.Set("body", body)
		c.Next()
	}
}

func IsExcelFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("students")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, response.StudentResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		if fileExt := strings.ToLower(file.Filename[strings.LastIndex(file.Filename, ".")+1:]); fileExt != "xlsx" {
			c.AbortWithStatusJSON(http.StatusBadRequest, response.StudentResponse{
				Success: false,
				Message: "file uploaded is not xlsx file",
			})
			return
		}
		c.Next()
	}
}

func GetIDParam() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, response.StudentResponse{
				Success: false,
				Message: "path variables is not valid, please use a number",
			})
			return
		}
		c.Set("studentID", uint(id))
		c.Next()
	}
}


