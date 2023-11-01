package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
)

func ValidateHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		var headerToken = c.Request.Header.Get("Authorization")
		var bearer = strings.HasPrefix(headerToken, "Bearer")

		if !bearer {
			c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
				Success: false,
				Message: "Bearer token is not valid",
			})
			return
		}

		stringToken := strings.Split(headerToken, " ")[1]

		c.Set("jwt", stringToken)
		c.Next()
	}
}
