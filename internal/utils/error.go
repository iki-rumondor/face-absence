package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
)

func HandleError(c *gin.Context, err error) {
	res := err.(*response.Error)

	switch res.Code {
	case 400:
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Success: false,
			Message: res.Message,
		})
		return
	case 401:
		c.AbortWithStatusJSON(http.StatusUnauthorized, response.FailedResponse{
			Success: false,
			Message: res.Message,
		})
		return
	case 404:
		c.AbortWithStatusJSON(http.StatusNotFound, response.FailedResponse{
			Success: false,
			Message: res.Message,
		})
		return
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.FailedResponse{
			Success: false,
			Message: res.Message,
		})
		return
	}
}

func IsErrorType(err error) bool {
	_, ok := err.(*response.Error)
	return ok
}