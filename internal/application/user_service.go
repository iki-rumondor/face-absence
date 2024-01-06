package application

import (
	"fmt"
	"os"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
)

type UserService struct {
	Repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) UpdateAvatar(model *domain.User) error {
	user, err := s.Repo.FindUserByID(model.ID)
	if err != nil {
		return &response.Error{
			Code:    404,
			Message: "User is not found",
		}
	}

	
	if err := s.Repo.UpdateAvatar(model); err != nil {
		// Hapus File Di Folder
		if err := os.Remove("internal/assets/avatar/" + *model.Avatar); err != nil {
			fmt.Println(err.Error())
		}
		
		return &response.Error{
			Code:    500,
			Message: "Failed to update user: " + err.Error(),
		}
	}
	
	if user.Avatar != nil{
		if err := os.Remove("internal/assets/avatar/" + *user.Avatar); err != nil {
			fmt.Println(err.Error())
		}
	}

	return nil
}
