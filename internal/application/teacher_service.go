package application

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"gorm.io/gorm"
)

type TeacherService struct {
	Repo repository.TeacherRepository
}

func NewTeacherService(repo repository.TeacherRepository) *TeacherService {
	return &TeacherService{
		Repo: repo,
	}
}

func (s *TeacherService) CreateTeacher(request request.CreateTeacher) error {
	user := &domain.User{
		Uuid:     uuid.NewString(),
		Nama:     request.Nama,
		Email:    request.Email,
		Password: request.NIP,
		RoleID:   1,
	}

	user, err := s.Repo.SaveUser(user)

	if err != nil {
		return err
	}

	teacher := &domain.Teacher{
		Nip:    request.NIP,
		JK:     request.JK,
		UserID: user.ID,
	}

	if _, err := s.Repo.CreateTeacher(teacher); err != nil {
		s.Repo.DeleteUser(teacher.UserID)
		return err
	}

	return nil

}

func (s *TeacherService) GetTeachers() (*[]domain.Teacher, error) {

	teachers, err := s.Repo.FindTeachers()

	if err != nil {
		return nil, err
	}

	return teachers, nil

}

func (s *TeacherService) GetTeacher(uuid string) (*domain.Teacher, error) {

	teacher, err := s.Repo.FindTeacher(uuid)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("teacher with uuid %s is not found", uuid)
	}

	if err != nil {
		return nil, err
	}

	return teacher, nil

}

func (s *TeacherService) UpdateTeacher(request *request.UpdateTeacher) (*domain.Teacher, error) {

	user := &domain.User{
		Uuid: request.Uuid,
		Nama: request.Nama,
	}

	teacher := &domain.Teacher{
		Nip: request.NIP,
		JK:  request.JK,
	}

	result, err := s.Repo.UpdateTeacher(user, teacher)

	if err != nil {
		return nil, fmt.Errorf("failed to update teachers: %s", err.Error())
	}

	return result, nil

}

func (s *TeacherService) DeleteTeacherByUuid(uuid string) error {

	user, err := s.Repo.FindUserByUuid(uuid)
	if err != nil {
		return fmt.Errorf("teacher with uuid %s is not found", uuid)
	}

	if err := s.Repo.DeleteWithUser(user); err != nil {
		return fmt.Errorf("failed to delete teacher: %s", err.Error())
	}

	return nil

}
