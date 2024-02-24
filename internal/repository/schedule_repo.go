package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type ScheduleRepository interface {
	FindSchedulePagination(*domain.Pagination) (*domain.Pagination, error)
	CreateSchedule(*domain.Schedule) error
	FindSchedules() (*[]domain.Schedule, error)
	FindSchedulesByClass(classID uint) (*[]domain.Schedule, error)
	FindScheduleByUuid(string) (*domain.Schedule, error)
	FindAbsenceByDate(scheduleID uint, date string) (*[]domain.Absence, error)
	NewFindAbsenceByDate(scheduleID uint, date string) (*[]domain.Student, error)
	UpdateSchedule(*domain.Schedule) error
	DeleteSchedule(*domain.Schedule) error

	FindStudentsByClassID(classID uint) (*[]domain.Student, error)
	FindTeacherByUserID(userID uint) (*domain.Teacher, error)
	FindUserByID(ID uint) (*domain.User, error)
}
