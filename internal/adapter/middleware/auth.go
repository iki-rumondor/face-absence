package middleware

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
)

func ValidateLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body request.Login
		if err := c.BindJSON(&body); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
				Message: err.Error(),
			})
			return
		}

		if _, err := govalidator.ValidateStruct(&body); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
				Message: err.Error(),
			})
			return
		}

		c.Set("login", body)
		c.Next()
	}
}

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body request.JWT
		if err := c.BindJSON(&body); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
				Message: err.Error(),
			})
			return
		}

		if _, err := govalidator.ValidateStruct(&body); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
				Message: err.Error(),
			})
			return
		}

		c.Set("jwt", body.Token)
		c.Next()
	}
}
