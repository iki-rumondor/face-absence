package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type ScheduleRepository interface {
	FindSchedules() (*[]domain.Schedule, error)
	FindScheduleByUuid(string) (*domain.Schedule, error)
	SaveSchedule(*domain.Schedule) error
	DeleteSchedule(*domain.Schedule) error
}
