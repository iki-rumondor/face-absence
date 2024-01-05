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
			Message: "Failed to get all school years: " + err.Error(),
		}
	}

	page := GeneratePages(urlPath, result)

	return page, nil

}

func (s *SchoolYearService) CreateSchoolYear(model *domain.SchoolYear) error {

	if err := s.Repo.CreateSchoolYear(model); err != nil {
		return &response.Error{
			Code:    500,
			Message: "SchoolYear was not created successfully: " + err.Error(),
		}
	}

	return nil
}

func (s *SchoolYearService) GetAllSchoolYears() (*[]response.SchoolYearResponse, error) {

	result, err := s.Repo.FindSchoolYears()

	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Failed: " + err.Error(),
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

func (s *SchoolYearService) GetSchoolYear(uuid string) (*response.SchoolYearResponse, error) {

	result, err := s.Repo.FindSchoolYearByUuid(uuid)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.Error{
				Code:    404,
				Message: fmt.Sprintf("SchoolYear with uuid %s is not found", uuid),
			}
		}
		return nil, &response.Error{
			Code:    500,
			Message: "Failed: " + err.Error(),
		}
	}

	res := response.SchoolYearResponse{
		Uuid:      result.Uuid,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	return &res, nil
}

func (s *SchoolYearService) UpdateSchoolYear(model *domain.SchoolYear) error {

	if err := s.Repo.UpdateSchoolYear(model); err != nil {
		return &response.Error{
			Code:    500,
			Message: "SchoolYear was not updated successfully: " + err.Error(),
		}
	}

	return nil
}

func (s *SchoolYearService) DeleteSchoolYear(SchoolYear *domain.SchoolYear) error {

	if err := s.Repo.DeleteSchoolYear(SchoolYear); err != nil {
		return &response.Error{
			Code:    500,
			Message: "SchoolYear was not deleted successfully: " + err.Error(),
		}
	}

	return nil
}
