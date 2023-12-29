package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type ClassRepository interface {
	FindClasses() (*[]domain.Class, error)
	FindClassByUuid(string) (*domain.Class, error)
	SaveClass(*domain.Class) error
	DeleteClass(*domain.Class) error
}
