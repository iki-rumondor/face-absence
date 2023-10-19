package middleware

import (
	"net/http"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
)

func ValidateBodyJSON() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body request.Student

		if err := c.BindJSON(&body); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, response.Student{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		if _, err := govalidator.ValidateStruct(&body); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, response.Student{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		c.Next()
	}
}

func IsExcelFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("students")
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Student{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		if fileExt := strings.ToLower(file.Filename[strings.LastIndex(file.Filename, ".")+1:]); fileExt != "xlsx" {
			c.JSON(http.StatusBadRequest, response.Student{
				Success: false,
				Message: "file uploaded is not xlsx file",
			})
			return
		}
		c.Next()
	}
}
