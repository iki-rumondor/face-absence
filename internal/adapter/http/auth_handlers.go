package customHTTP

import (
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/application"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
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
		Username: body.Username,
		Password: body.Password,
	}

	jwt, err := h.Service.VerifyUser(body.Role, &user)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.JWT{
		Token: jwt,
	})
}

func (h *AuthHandlers) VerifyToken(c *gin.Context) {

	userID := c.GetUint("user_id")
	userRole := c.GetString("user_role")

	user, err := h.Service.GetUserByID(userID)
	if err != nil {
		utils.HandleError(c, err)
	}

	res := response.VerifyTokenResponse{
		Nama:      user.Nama,
		Username:  user.Username,
		Avatar:    fmt.Sprintf("/public/avatar/%s", *user.Avatar),
		Role:      userRole,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Token valid",
		Data:    &res,
	})
}
