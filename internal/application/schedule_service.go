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

func (s *ScheduleService) GetAllSchedules() (*[]domain.Schedule, error) {

	result, err := s.Repo.FindSchedules()
	if err != nil {
		return nil, INTERNAL_ERROR
	}

	return result, nil
}

// func (s *ScheduleService) GetScheduleStudentNow(userID uint, scheduleUuid string) (*domain.Schedule, *domain.Absence, error) {

// 	student, err := s.Repo.FindStudentByUserID(userID)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	schedule, err := s.GetSchedule(scheduleUuid)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	absence, err := s.Repo.FindStudentAbsenceByScheduleID(student.ID, schedule.ID)
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return schedule, nil, nil
// 		}
// 		return nil, nil, INTERNAL_ERROR
// 	}

// 	return schedule, absence, nil
// }

func (s *ScheduleService) GetTeacherSchedules(userID uint) (*[]domain.Schedule, error) {

	teacher, err := s.Repo.FindTeacherByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.Error{
				Code:    404,
				Message: "Guru tidak ditemukan",
			}
		}
		return nil, INTERNAL_ERROR
	}

	var schedules []domain.Schedule

	for _, item := range *teacher.Subjects{
		schedules = append(schedules, *item.Schedules...)
	}

	if err != nil {
		return nil, INTERNAL_ERROR
	}

	return &schedules, nil
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
