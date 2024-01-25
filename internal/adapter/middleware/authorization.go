package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
)

var (
	UNAUTHERROR = &response.Error{
		Code:    403,
		Message: "Hak akses dibatasi",
	}
)

func IsRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mc := c.MustGet("map_claims")
		mapClaims := mc.(jwt.MapClaims)

		roleJwt := mapClaims["role"].(string)
		if roleJwt != role {
			utils.HandleError(c, UNAUTHERROR)
			return
		}
		c.Next()
	}
}
