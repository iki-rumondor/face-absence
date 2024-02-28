package application

import (
	"errors"
	"fmt"
	"os"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"gorm.io/gorm"
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

func (s *UserService) GetDashboardData() (map[string]int64, error) {
	res, err := s.Repo.CountStudentsTeachersAdmins()
	if err != nil {
		return nil, INTERNAL_ERROR
	}

	res["admin_woman"] = res["admin"] - res["admin_man"]
	res["student_woman"] = res["student"] - res["student_man"]
	res["teacher_woman"] = res["teacher"] - res["teacher_man"]

	return res, nil
}

func (s *UserService) UpdatePassword(uuid string, req request.ChangePassword) error {

	if req.NewPassword != req.ConfirmPassword {
		return &response.Error{
			Code:    400,
			Message: "Password Tidak Sama",
		}
	}

	teacher, err := s.Repo.FindTeacherByUuid(uuid)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return &response.Error{
				Code:    404,
				Message: "Guru Tidak Ditemukan",
			}
		}
		return INTERNAL_ERROR
	}

	model := domain.User{
		ID:       teacher.UserID,
		Password: req.NewPassword,
	}

	if err := s.Repo.Update(&model); err != nil {
		return INTERNAL_ERROR
	}

	return nil
}

func (s *UserService) GetAllUsers() (*[]domain.User, error) {

	user, err := s.Repo.FindUsers()
	if err != nil {
		return nil, INTERNAL_ERROR
	}

	return user, nil
}
