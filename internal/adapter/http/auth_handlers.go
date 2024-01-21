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
	if userID == 0 {
		utils.HandleError(c, &response.Error{
			Code:    400,
			Message: "HandleError: ID User tidak dapat ditemukan",
		})
	}

	user, err := h.Service.GetUserByID(userID)
	if err != nil {
		utils.HandleError(c, err)
	}

	avatar := fmt.Sprintf("/public/avatar/%s", *user.Avatar)
	res := response.UserData{
		Nama:      user.Nama,
		Username:  user.Username,
		Avatar:    &avatar,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Success: true,
		Message: "Token valid",
		Data:    &res,
	})
}
