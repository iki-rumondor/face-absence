package application

import (
	"github.com/google/uuid"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
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

	user, err := s.Repo.CreateUser(user)

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
