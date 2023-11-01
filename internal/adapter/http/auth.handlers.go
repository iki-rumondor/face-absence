package customHTTP

import (
	"net/http"

	"github.com/asaskevich/govalidator"
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
	var body request.Login
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.FailedResponse{
			Message: err.Error(),
		})
		return
	}

	user := domain.User{
		Email:    body.Email,
		Password: body.Password,
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

	mapClaims, err := h.Service.VerifyToken(jwt)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.FailedResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "your JWT token is authenticated and good to go",
		Data:    mapClaims,
	})
}
