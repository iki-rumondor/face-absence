package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type AbsenceRepository interface {
	FindUserByID(uint) (*domain.User, error)
	FindScheduleByID(uint) (*domain.Schedule, error)
	CreateAbsence(*domain.Absence) error
}
