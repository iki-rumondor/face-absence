package customHTTP

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/application"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
)

type AuthHandlers struct {
	Service *application.AuthService
}

func NewAuthHandler(service *application.AuthService) *AuthHandlers {
	return &AuthHandlers{
		Service: service,
	}
}

func (h *AuthHandlers) Login(c *gin.Context) {
	body, ok := c.Get("login")

	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.FailedResponse{
			Success: false,
			Message: "something went wrong at user service",
		})
		return
	}

	var login = body.(request.Login)

	user := domain.User{
		Email:    login.Email,
		Password: login.Password,
	}

	jwt, err := h.Service.VerifyUser(&user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.FailedResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.JWT{
		Token: jwt,
	})
}

func (h *AuthHandlers) VerifyToken(c *gin.Context) {

	jwt := c.GetString("jwt")

	if err := h.Service.VerifyToken(jwt); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.FailedResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "your JWT token is authenticated and good to go",
	})
}
