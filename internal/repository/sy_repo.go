package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type SchoolYearRepository interface {
	CreateSchoolYear(*domain.SchoolYear) error
	FindSchoolYears() (*[]domain.SchoolYear, error)
	FindSchoolYearByUuid(string) (*domain.SchoolYear, error)
	UpdateSchoolYear(*domain.SchoolYear) error
	DeleteSchoolYear(*domain.SchoolYear) error
}
