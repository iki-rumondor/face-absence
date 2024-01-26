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

type SubjectService struct {
	Repo repository.SubjectRepository
}

func NewSubjectService(repo repository.SubjectRepository) *SubjectService {
	return &SubjectService{
		Repo: repo,
	}
}

func (s *SubjectService) CreateSubject(body *request.CreateSubject) error {
	var teachers []domain.Teacher

	for i, item := range body.TeachersUuid {
		teacher, err := s.Repo.FindTeacherByUuid(string(item[i]))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return &response.Error{
					Code:    404,
					Message: fmt.Sprintf("Guru dengan uuid %s tidak ditemukan", string(item[i])),
				}
			}
			return INTERNAL_ERROR
		}

		teachers = append(teachers, *teacher)
	}

	subject := domain.Subject{
		Uuid:     uuid.NewString(),
		Name:     body.Name,
		Teachers: teachers,
	}

	if err := s.Repo.CreateSubject(&subject); err != nil {
		return INTERNAL_ERROR
	}

	return nil
}

func (s *SubjectService) SubjectPagination(urlPath string, pagination *domain.Pagination) (*domain.Pagination, error) {

	result, err := s.Repo.FindSubjectPagination(pagination)
	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		}
	}

	page := GeneratePages(urlPath, result)

	return page, nil

}

func (s *SubjectService) GetAllSubjects() (*[]domain.Subject, error) {

	result, err := s.Repo.FindSubjects()

	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		}
	}

	return result, nil
}

func (s *SubjectService) GetSubject(uuid string) (*domain.Subject, error) {

	result, err := s.Repo.FindSubjectByUuid(uuid)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.Error{
				Code:    404,
				Message: fmt.Sprintf("Mata pelajaran dengan uuid %s tidak ditemukan", uuid),
			}
		}
		return nil, &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		}
	}

	return result, nil
}

func (s *SubjectService) UpdateSubject(subjectUuid string, body *request.UpdateSubject) error {

	res, err := s.GetSubject(subjectUuid)
	if err != nil {
		return err
	}

	var teachers []domain.Teacher

	for i, item := range body.TeachersUuid {
		teacher, err := s.Repo.FindTeacherByUuid(string(item[i]))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return &response.Error{
					Code:    404,
					Message: fmt.Sprintf("Guru dengan uuid %s tidak ditemukan", string(item[i])),
				}
			}
			return INTERNAL_ERROR
		}

		teachers = append(teachers, *teacher)
	}

	model := domain.Subject{
		ID:       res.ID,
		Name:     body.Name,
		Teachers: teachers,
	}

	if err := s.Repo.UpdateSubject(&model); err != nil {
		return INTERNAL_ERROR
	}

	return nil
}

func (s *SubjectService) DeleteSubject(Subject *domain.Subject) error {

	if err := s.Repo.DeleteSubject(Subject); err != nil {

		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return &response.Error{
				Code:    403,
				Message: "Data ini tidak dapat dihapus karena berelasi dengan data lain",
			}
		}
		return &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		}
	}

	return nil
}
