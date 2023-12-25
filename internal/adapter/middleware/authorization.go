package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
)

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		mc := c.MustGet("map_claims")
		mapClaims := mc.(jwt.MapClaims)

		role := mapClaims["role"].(string)
		if role != "ADMIN" {
			c.AbortWithStatusJSON(http.StatusInternalServerError, response.FailedResponse{
				Success: false,
				Message: "access denied due to invalid credentials",
			})
			return
		}
		c.Next()
	}
}

func IsTeacher() gin.HandlerFunc {
	return func(c *gin.Context) {
		mc := c.MustGet("map_claims")
		mapClaims := mc.(jwt.MapClaims)

		role := mapClaims["role"].(string)
		if role != "GURU" {
			c.AbortWithStatusJSON(http.StatusInternalServerError, response.FailedResponse{
				Success: false,
				Message: "access denied due to invalid credentials",
			})
			return
		}
		c.Next()
	}
}

func IsStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		mc := c.MustGet("map_claims")
		mapClaims := mc.(jwt.MapClaims)

		role := mapClaims["role"].(string)
		if role != "SANTRI" {
			c.AbortWithStatusJSON(http.StatusInternalServerError, response.FailedResponse{
				Success: false,
				Message: "access denied due to invalid credentials",
			})
			return
		}
		c.Next()
	}
}
