package application

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
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

	if user, _ := s.Repo.FindUserByUsername(request.Username); user != nil {
		return &response.Error{
			Code:    404,
			Message: "Username has already been taken",
		}
	}

	if user, _ := s.Repo.FindTeacherByColumn("nip", request.Nip); user != nil {
		return &response.Error{
			Code:    404,
			Message: "Nip has already been taken",
		}
	}

	user := &domain.User{
		Nama:     request.Nama,
		Username: request.Username,
		Password: request.Password,
	}

	teacher := &domain.Teacher{
		Uuid:          uuid.NewString(),
		Nuptk:         request.Nuptk,
		Nip:           request.Nip,
		StatusPegawai: request.StatusPegawai,
		JK:            request.JK,
		TempatLahir:   request.TempatLahir,
		TanggalLahir:  request.TanggalLahir,
		NoHp:          request.NoHp,
		Jabatan:       request.Jabatan,
		TotalJtm:      request.TotalJtm,
		Alamat:        request.Alamat,
	}

	if err := s.Repo.CreateTeacherUser(teacher, user); err != nil {
		return &response.Error{
			Code:    500,
			Message: "Failed to create user: " + err.Error(),
		}
	}

	return nil
}

func (s *TeacherService) TeachersPagination(urlPath string, pagination *domain.Pagination) (*domain.Pagination, error) {

	result, err := s.Repo.FindTeachersPagination(pagination)
	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Failed to get all users: " + err.Error(),
		}
	}

	page := GeneratePages(urlPath, result)

	return page, nil

}

func (s *TeacherService) GetTeachers() (*[]domain.Teacher, error) {

	teachers, err := s.Repo.FindTeachers()

	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Failed to get all teachers" + err.Error(),
		}
	}

	return teachers, nil

}

func (s *TeacherService) GetTeacher(uuid string) (*domain.Teacher, error) {

	teacher, err := s.Repo.FindTeacherByUuid(uuid)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, &response.Error{
			Code:    404,
			Message: fmt.Sprintf("Teacher with uuid %s is not found", uuid),
		}
	}

	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Failed to get teacher: " + err.Error(),
		}
	}

	return teacher, nil

}

func (s *TeacherService) UpdateTeacher(request *request.UpdateTeacher) error {

	teacherInDB, err := s.Repo.FindTeacherByUuid(request.Uuid)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &response.Error{
			Code:    404,
			Message: "Teacher with uuid is not found",
		}
	}

	if err != nil {
		return &response.Error{
			Code:    500,
			Message: "Failed to get teacher: " + err.Error(),
		}
	}

	user := &domain.User{
		Nama:     request.Nama,
		Username: request.Username,
	}

	teacher := &domain.Teacher{
		Uuid:          request.Uuid,
		Nuptk:         request.Nuptk,
		Nip:           request.Nip,
		StatusPegawai: request.StatusPegawai,
		JK:            request.JK,
		TempatLahir:   request.TempatLahir,
		TanggalLahir:  request.TanggalLahir,
		NoHp:          request.NoHp,
		Jabatan:       request.Jabatan,
		TotalJtm:      request.TotalJtm,
		Alamat:        request.Alamat,
		UserID:        teacherInDB.UserID,
	}

	if err := s.Repo.UpdateTeacherUser(teacher, user); err != nil {
		return &response.Error{
			Code:    500,
			Message: "Failed to update teacher: " + err.Error(),
		}
	}

	return nil

}

func (s *TeacherService) DeleteTeacher(uuid string) error {

	teacherInDB, err := s.Repo.FindTeacherByUuid(uuid)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &response.Error{
			Code:    404,
			Message: "Teacher with uuid is not found",
		}
	}

	if err != nil {
		return &response.Error{
			Code:    500,
			Message: "Failed to get teacher: " + err.Error(),
		}
	}

	if err := s.Repo.DeleteTeacherUser(teacherInDB.UserID); err != nil {
		return &response.Error{
			Code:    500,
			Message: "Failed to delete teacher: " + err.Error(),
		}
	}

	return nil

}
