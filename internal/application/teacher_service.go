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
			Message: "Username sudah terdaftar",
		}
	}

	if user, _ := s.Repo.FindTeacherByColumn("nip", request.Nip); user != nil {
		return &response.Error{
			Code:    404,
			Message: "Nip sudah terdaftar",
		}
	}

	user := &domain.User{
		Nama:     request.Nama,
		Username: request.Username,
		Password: request.Username,
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
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		}
	}

	return nil
}

func (s *TeacherService) TeachersPagination(urlPath string, pagination *domain.Pagination) (*domain.Pagination, error) {

	result, err := s.Repo.FindTeachersPagination(pagination)
	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
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
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		}
	}

	return teachers, nil

}

func (s *TeacherService) GetTeacher(uuid string) (*domain.Teacher, error) {

	teacher, err := s.Repo.FindTeacherByUuid(uuid)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, &response.Error{
			Code:    404,
			Message: fmt.Sprintf("Guru dengan uuid %s tidak ditemukan", uuid),
		}
	}

	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		}
	}

	return teacher, nil

}

func (s *TeacherService) UpdateTeacher(request *request.UpdateTeacher) error {

	teacherInDB, err := s.Repo.FindTeacherByUuid(request.Uuid)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &response.Error{
			Code:    404,
			Message: fmt.Sprintf("Guru dengan uuid %s tidak ditemukan", request.Uuid),
		}
	}

	if err != nil {
		return &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
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
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		}
	}

	return nil

}

func (s *TeacherService) DeleteTeacher(uuid string) error {

	if err := s.Repo.DeleteTeacher(uuid); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &response.Error{
				Code:    404,
				Message: fmt.Sprintf("Guru dengan uuid %s tidak ditemukan", uuid),
			}
		}

		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return &response.Error{
				Code:    500,
				Message: "Data ini tidak dapat dihapus karena berelasi dengan data lain",
			}
		}
		return INTERNAL_ERROR
	}

	return nil

}
