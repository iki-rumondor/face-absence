package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type ScheduleRepository interface {
	FindSchedulePagination(*domain.Pagination) (*domain.Pagination, error)
	CreateSchedule(*domain.Schedule) error
	FindSchedules() (*[]domain.Schedule, error)
	FindSchedulesByClass(classID uint) (*[]domain.Schedule, error)
	FindScheduleByUuid(string) (*domain.Schedule, error)
	UpdateSchedule(*domain.Schedule) error
	DeleteSchedule(*domain.Schedule) error

	FindStudentByUserID(userID uint) (*domain.Student, error)
	FindUserByID(ID uint) (*domain.User, error)
	FindStudentAbsenceByScheduleID(studentID, scheduleID uint) (*domain.Absence, error)
}
