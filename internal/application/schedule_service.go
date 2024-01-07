package application

import (
	"errors"

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

func (s *ScheduleService) SchedulePagination(urlPath string, pagination *domain.Pagination) (*domain.Pagination, error) {

	result, err := s.Repo.FindSchedulePagination(pagination)
	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Gagal mendapatkan jadwal",
		}
	}

	page := GeneratePages(urlPath, result)

	return page, nil

}

func (s *ScheduleService) CreateSchedule(model *domain.Schedule) error {

	if err := s.Repo.CreateSchedule(model); err != nil {
		return &response.Error{
			Code:    500,
			Message: "Gagal menambahkan jadwal",
		}
	}

	return nil
}

func (s *ScheduleService) GetAllSchedules() (*[]response.ScheduleResponse, error) {

	result, err := s.Repo.FindSchedules()

	if err != nil {
		return nil, &response.Error{
			Code:    500,
			Message: "Gagal mendapatkan seluruh jadwal",
		}
	}

	var resp []response.ScheduleResponse

	for _, res := range *result {
		resp = append(resp, response.ScheduleResponse{
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

func (s *ScheduleService) GetSchedule(uuid string) (*domain.Schedule, error) {

	result, err := s.Repo.FindScheduleByUuid(uuid)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.Error{
				Code:    404,
				Message: "Jadwal tidak ditemukan",
			}
		}
		return nil, &response.Error{
			Code:    500,
			Message: "Terjadi kesalahan dalam mengambil jadwal",
		}
	}

	return result, nil
}

func (s *ScheduleService) UpdateSchedule(model *domain.Schedule) error {

	if err := s.Repo.UpdateSchedule(model); err != nil {
		if utils.IsErrorType(err) {
			return err
		}
		return &response.Error{
			Code:    500,
			Message: "Jadwal gagal diupdate",
		}
	}

	return nil
}

func (s *ScheduleService) DeleteSchedule(Schedule *domain.Schedule) error {

	if err := s.Repo.DeleteSchedule(Schedule); err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return &response.Error{
				Code:    403,
				Message: "Data ini tidak dapat dihapus karena berelasi dengan data lain",
			}
		}
		return &response.Error{
			Code:    500,
			Message: "Jadwal gagal dihapus",
		}
	}

	return nil
}
