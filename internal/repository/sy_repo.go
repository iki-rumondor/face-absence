package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type SchoolYearRepository interface {
	FindSchoolYears() (*[]domain.SchoolYear, error)
	FindSchoolYearByUuid(string) (*domain.SchoolYear, error)
	SaveSchoolYear(*domain.SchoolYear) error
	DeleteSchoolYear(*domain.SchoolYear) error
}
