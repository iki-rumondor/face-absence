package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type ScheduleRepository interface {
	CreateSchedule(*domain.Schedule) error
	FindSchedules() (*[]domain.Schedule, error)
	FindScheduleByUuid(string) (*domain.Schedule, error)
	UpdateSchedule(*domain.Schedule) error
	DeleteSchedule(*domain.Schedule) error
}
