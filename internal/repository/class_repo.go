package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type ClassRepository interface {
	CreateClass(*domain.Class) error
	FindClasses() (*[]domain.Class, error)
	FindClassByUuid(string) (*domain.Class, error)
	UpdateClass(*domain.Class) error
	DeleteClass(*domain.Class) error
}
