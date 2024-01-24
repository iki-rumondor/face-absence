package application

import (
	"errors"
	"fmt"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
	"gorm.io/gorm"
)

type ClassService struct {
	Repo repository.ClassRepository
}

func NewClassService(repo repository.ClassRepository) *ClassService {
	return &ClassService{
		Repo: repo,
	}
}

func (s *ClassService) CreateClass(class *domain.Class) error {

	if err := s.Repo.CreateClass(class); err != nil {
		if utils.IsErrorType(err) {
			return err
		}

		return &response.Error{
			Code:    500,
			Message: "Gagal menambahkan kelas, silahkan hubungi developper",
		}
	}

	return nil
}

func (s *ClassService) GetClassOptions() (*[]response.ClassOption, error) {

	classes, err := s.Repo.FindClasses()

	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Gagal menemukan kelas, silahkan hubungi developper",
		}
	}

	var res []response.ClassOption

	for _, class := range *classes {
		res = append(res, response.ClassOption{
			Uuid: class.Uuid,
			Name: class.Name,
		})
	}

	return &res, nil
}

func (s *ClassService) GetAllClasses() (*[]response.ClassResponse, error) {

	classes, err := s.Repo.FindClasses()

	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Gagal menemukan kelas, silahkan hubungi developper",
		}
	}

	var res []response.ClassResponse

	for _, class := range *classes {
		res = append(res, response.ClassResponse{
			Uuid: class.Uuid,
			Name: class.Name,
			Teacher: &response.Teacher{
				Uuid:          class.Teacher.Uuid,
				JK:            class.Teacher.JK,
				Nip:           class.Teacher.Nip,
				Nuptk:         class.Teacher.Nuptk,
				StatusPegawai: class.Teacher.StatusPegawai,
				TempatLahir:   class.Teacher.TempatLahir,
				TanggalLahir:  class.Teacher.TanggalLahir,
				NoHp:          class.Teacher.NoHp,
				Jabatan:       class.Teacher.Jabatan,
				TotalJtm:      class.Teacher.TotalJtm,
				Alamat:        class.Teacher.Alamat,
				User: &response.UserData{
					Nama:      class.Teacher.User.Nama,
					Username:  class.Teacher.User.Username,
					Avatar:    class.Teacher.User.Avatar,
					CreatedAt: class.Teacher.User.CreatedAt,
					UpdatedAt: class.Teacher.User.UpdatedAt,
				},
				CreatedAt: class.Teacher.CreatedAt,
				UpdatedAt: class.Teacher.UpdatedAt,
			},
			CreatedAt: class.CreatedAt,
			UpdatedAt: class.UpdatedAt,
		})
	}

	return &res, nil
}

func (s *ClassService) ClassPagination(urlPath string, pagination *domain.Pagination) (*domain.Pagination, error) {

	result, err := s.Repo.FindClassPagination(pagination)
	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Gagal menemukan kelas, silahkan hubungi developper",
		}
	}

	page := GeneratePages(urlPath, result)

	return page, nil

}

func (s *ClassService) GetClass(uuid string) (*domain.Class, error) {

	class, err := s.Repo.FindClassByUuid(uuid)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.Error{
				Code:    404,
				Message: fmt.Sprintf("Kelas dengan uuid %s tidak ditemukan", uuid),
			}
		}
		return nil, &response.Error{
			Code:    500,
			Message: "Gagal menemukan kelas, silahkan hubungi developper",
		}
	}

	return class, nil
}

func (s *ClassService) UpdateClass(class *domain.Class) error {

	if err := s.Repo.UpdateClass(class); err != nil {
		if utils.IsErrorType(err) {
			return err
		}

		return &response.Error{
			Code:    500,
			Message: "Gagal mengupdate kelas, silahkan hubungi developper",
		}
	}

	return nil
}

func (s *ClassService) DeleteClass(class *domain.Class) error {

	if err := s.Repo.DeleteClass(class); err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return &response.Error{
				Code:    403,
				Message: "Data ini tidak dapat dihapus karena berelasi dengan data lain",
			}
		}
		return &response.Error{
			Code:    500,
			Message: "Gagal menghapus kelas, silahkan hubungi developper",
		}
	}

	return nil
}

func (s *ClassService) CreateClassPDF() ([]byte, error) {

	classes, err := s.GetAllClasses()
	if err != nil {
		return nil, err
	}

	if len(*classes) == 0 {
		return nil, &response.Error{
			Code:    404,
			Message: "ServiceError: Data Kelas Masih Kosong",
		}
	}

	var data []*request.ClassPDFData

	for _, item := range *classes {
		data = append(data, &request.ClassPDFData{
			Name:        item.Name,
			TeacherName: item.Teacher.User.Nama,
			CreatedAt:   item.CreatedAt.Format("02 Januari 2006"),
		})
	}

	pdfData, err := s.Repo.GetClassPDF(data)
	if err != nil {
		return nil, INTERNAL_ERROR
	}

	return pdfData, nil
}
