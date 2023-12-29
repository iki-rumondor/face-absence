package application

import (
	"errors"
	"fmt"

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

func (s *SubjectService) CreateSubject(model *domain.Subject) error {

	if err := s.Repo.SaveSubject(model); err != nil {
		return &response.Error{
			Code:    500,
			Message: "Subject was not created successfully: " + err.Error(),
		}
	}

	return nil
}

func (s *SubjectService) GetAllSubjects() (*[]response.SubjectResponse, error) {

	result, err := s.Repo.FindSubjects()

	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Failed: " + err.Error(),
		}
	}

	var resp []response.SubjectResponse

	for _, res := range *result {
		resp = append(resp, response.SubjectResponse{
			ID:        res.ID,
			Uuid:      res.Uuid,
			Name:      res.Name,
			CreatedAt: res.CreatedAt,
			UpdatedAt: res.UpdatedAt,
		})
	}

	return &resp, nil
}

func (s *SubjectService) GetSubject(uuid string) (*response.SubjectResponse, error) {

	result, err := s.Repo.FindSubjectByUuid(uuid)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.Error{
				Code:    404,
				Message: fmt.Sprintf("Subject with uuid %s is not found", uuid),
			}
		}
		return nil, &response.Error{
			Code:    500,
			Message: "Failed: " + err.Error(),
		}
	}

	res := response.SubjectResponse{
		ID:        result.ID,
		Uuid:      result.Uuid,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	return &res, nil
}

func (s *SubjectService) UpdateSubject(model *domain.Subject) error {

	if err := s.Repo.SaveSubject(model); err != nil {
		return &response.Error{
			Code:    500,
			Message: "Subject was not updated successfully: " + err.Error(),
		}
	}

	return nil
}

func (s *SubjectService) DeleteSubject(Subject *domain.Subject) error {

	if err := s.Repo.DeleteSubject(Subject); err != nil {
		return &response.Error{
			Code:    500,
			Message: "Subject was not deleted successfully: " + err.Error(),
		}
	}

	return nil
}
