package application

import (
	"errors"
	"fmt"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"gorm.io/gorm"
)

type SchoolYearService struct {
	Repo repository.SchoolYearRepository
}

func NewSchoolYearService(repo repository.SchoolYearRepository) *SchoolYearService {
	return &SchoolYearService{
		Repo: repo,
	}
}

func (s *SchoolYearService) SchoolYearPagination(urlPath string, pagination *domain.Pagination) (*domain.Pagination, error) {

	result, err := s.Repo.FindSchoolYearPagination(pagination)
	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		}
	}

	page := GeneratePages(urlPath, result)

	return page, nil

}

func (s *SchoolYearService) CreateSchoolYear(model *domain.SchoolYear) error {

	if err := s.Repo.CreateSchoolYear(model); err != nil {
		return &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		}
	}

	return nil
}

func (s *SchoolYearService) GetAllSchoolYears() (*[]response.SchoolYearResponse, error) {

	result, err := s.Repo.FindSchoolYears()

	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		}
	}

	var resp []response.SchoolYearResponse

	for _, res := range *result {
		resp = append(resp, response.SchoolYearResponse{
			Uuid:      res.Uuid,
			Name:      res.Name,
			CreatedAt: res.CreatedAt,
			UpdatedAt: res.UpdatedAt,
		})
	}

	return &resp, nil
}

func (s *SchoolYearService) GetSchoolYear(uuid string) (*domain.SchoolYear, error) {

	result, err := s.Repo.FindSchoolYearByUuid(uuid)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.Error{
				Code:    404,
				Message: fmt.Sprintf("Tahun pelajaran dengan uuid %s tidak ditemukan", uuid),
			}
		}
		return nil, &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		}
	}

	return result, nil
}

func (s *SchoolYearService) UpdateSchoolYear(model *domain.SchoolYear) error {

	if err := s.Repo.UpdateSchoolYear(model); err != nil {
		return &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan sistem, silahkan hubungi developper",
		}
	}

	return nil
}

func (s *SchoolYearService) DeleteSchoolYear(SchoolYear *domain.SchoolYear) error {

	if err := s.Repo.DeleteSchoolYear(SchoolYear); err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return &response.Error{
				Code:    500,
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
