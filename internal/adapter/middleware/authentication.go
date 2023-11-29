package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
	"gorm.io/gorm"
)

func IsValidJWT(db *gorm.DB) gin.HandlerFunc {
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

		jwt := strings.Split(headerToken, " ")[1]
		
		mapClaims, err := utils.VerifyToken(jwt)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.FailedResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		id := uint(mapClaims["id"].(float64))

		if err := db.First(&domain.User{}, "id = ?", id).Error; err != nil{
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.FailedResponse{
				Success: false,
				Message: "you are not registered in this system",
			})
			return
		}

		c.Set("map_claims", mapClaims)
		c.Next()

	}
}

