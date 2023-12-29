package application

import (
	"errors"
	"fmt"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
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
		return &response.Error{
			Code:    500,
			Message: "Failed to create class: " + err.Error(),
		}
	}

	return nil
}

func (s *ClassService) GetAllClasses() (*[]response.ClassResponse, error) {

	classes, err := s.Repo.FindClasses()

	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Failed to find classes: " + err.Error(),
		}
	}

	var res []response.ClassResponse

	for _, class := range *classes {
		res = append(res, response.ClassResponse{
			ID:        class.ID,
			Uuid:      class.Uuid,
			Name:      class.Name,
			TeacherID: class.TeacherID,
			CreatedAt: class.CreatedAt,
			UpdatedAt: class.UpdatedAt,
		})
	}

	return &res, nil
}

func (s *ClassService) GetClass(uuid string) (*response.ClassResponse, error) {

	class, err := s.Repo.FindClassByUuid(uuid)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.Error{
				Code:    404,
				Message: fmt.Sprintf("Class with uuid %s is not found", uuid),
			}
		}
		return nil, &response.Error{
			Code:    500,
			Message: "Failed to create class: " + err.Error(),
		}
	}

	res := response.ClassResponse{
		ID:        class.ID,
		Uuid:      class.Uuid,
		Name:      class.Name,
		TeacherID: class.TeacherID,
		CreatedAt: class.CreatedAt,
		UpdatedAt: class.UpdatedAt,
	}

	return &res, nil
}

func (s *ClassService) UpdateClass(class *domain.Class) error {

	if err := s.Repo.UpdateClass(class); err != nil {
		return &response.Error{
			Code:    500,
			Message: "Failed to update class: " + err.Error(),
		}
	}

	return nil
}

func (s *ClassService) DeleteClass(class *domain.Class) error {

	if err := s.Repo.DeleteClass(class); err != nil {
		return &response.Error{
			Code:    500,
			Message: "Failed to delete class: " + err.Error(),
		}
	}

	return nil
}
