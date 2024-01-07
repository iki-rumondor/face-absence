package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
)

func IsValidJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var headerToken = c.Request.Header.Get("Authorization")
		var bearer = strings.HasPrefix(headerToken, "Bearer")

		if !bearer {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.FailedResponse{
				Success: false,
				Message: "Bearer token is not valid",
			})
			return
		}

		jwt := strings.Split(headerToken, " ")[1]

		mapClaims, err := utils.VerifyToken(jwt)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.FailedResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		c.Set("map_claims", mapClaims)
		c.Next()

	}
}

func SetUserID() gin.HandlerFunc {
	return func(c *gin.Context) {
		mc := c.MustGet("map_claims")
		mapClaims := mc.(jwt.MapClaims)

		id, ok := mapClaims["id"].(float64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.FailedResponse{
				Success: false,
				Message: "Invalid JWT Token",
			})
			return
		}

		c.Set("user_id", uint(id))
		c.Next()

	}
}
