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
			Message: "User tidak dapat ditemukan",
		}
	}

	if err := s.Repo.UpdateAvatar(model); err != nil {
		// Hapus File Di Folder
		if err := os.Remove("internal/assets/avatar/" + *model.Avatar); err != nil {
			fmt.Println(err.Error())
		}

		return &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		}
	}

	if *user.Avatar != "default-avatar.jpg" {
		if err := os.Remove("internal/assets/avatar/" + *user.Avatar); err != nil {
			fmt.Println(err.Error())
		}
	}

	return nil
}

func (s *UserService) GetDashboardData() (map[string]int64 ,error) {
	res, err := s.Repo.CountStudentsTeachersAdmins()
	if err != nil {
		return nil, INTERNAL_ERROR
	}

	return res, nil
}
