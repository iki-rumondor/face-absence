package application

import (
	"errors"
	"fmt"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
	"gorm.io/gorm"
)

type ScheduleService struct {
	Repo repository.ScheduleRepository
}

func NewScheduleService(repo repository.ScheduleRepository) *ScheduleService {
	return &ScheduleService{
		Repo: repo,
	}
}

func (s *ScheduleService) CreateSchedule(model *domain.Schedule) error {

	if err := s.Repo.CreateSchedule(model); err != nil {
		if utils.IsErrorType(err) {
			return err
		}
		return &response.Error{
			Code:    500,
			Message: "Schedule was not created successfully: " + err.Error(),
		}
	}

	return nil
}

func (s *ScheduleService) GetAllSchedules() (*[]response.ScheduleResponse, error) {

	result, err := s.Repo.FindSchedules()

	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Failed: " + err.Error(),
		}
	}

	var resp []response.ScheduleResponse

	for _, res := range *result {
		resp = append(resp, response.ScheduleResponse{
			ID:           res.ID,
			Uuid:         res.Uuid,
			Name:         res.Name,
			Day:          res.Day,
			Start:        res.Start,
			End:          res.End,
			ClassID:      res.ClassID,
			SubjectID:    res.SubjectID,
			TeacherID:    res.TeacherID,
			SchoolYearID: res.SchoolYearID,
			CreatedAt:    res.CreatedAt,
			UpdatedAt:    res.UpdatedAt,
		})
	}

	return &resp, nil
}

func (s *ScheduleService) GetSchedule(uuid string) (*response.ScheduleResponse, error) {

	result, err := s.Repo.FindScheduleByUuid(uuid)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.Error{
				Code:    404,
				Message: fmt.Sprintf("Schedule with uuid %s is not found", uuid),
			}
		}
		return nil, &response.Error{
			Code:    500,
			Message: "Failed: " + err.Error(),
		}
	}

	res := response.ScheduleResponse{
		ID:           result.ID,
		Uuid:         result.Uuid,
		Name:         result.Name,
		Day:          result.Day,
		Start:        result.Start,
		End:          result.End,
		ClassID:      result.ClassID,
		SubjectID:    result.SubjectID,
		TeacherID:    result.TeacherID,
		SchoolYearID: result.SchoolYearID,
		CreatedAt:    result.CreatedAt,
		UpdatedAt:    result.UpdatedAt,
	}

	return &res, nil
}

func (s *ScheduleService) UpdateSchedule(model *domain.Schedule) error {

	if err := s.Repo.UpdateSchedule(model); err != nil {
		if utils.IsErrorType(err) {
			return err
		}
		return &response.Error{
			Code:    500,
			Message: "Schedule was not updated successfully: " + err.Error(),
		}
	}

	return nil
}

func (s *ScheduleService) DeleteSchedule(Schedule *domain.Schedule) error {

	if err := s.Repo.DeleteSchedule(Schedule); err != nil {
		return &response.Error{
			Code:    500,
			Message: "Schedule was not deleted successfully: " + err.Error(),
		}
	}

	return nil
}
