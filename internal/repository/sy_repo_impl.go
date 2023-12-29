package repository

import (
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type SchoolYearRepoImplementation struct {
	db *gorm.DB
}

func NewSchoolYearRepository(db *gorm.DB) SchoolYearRepository {
	return &SchoolYearRepoImplementation{
		db: db,
	}
}

func (r *SchoolYearRepoImplementation) SaveSchoolYear(model *domain.SchoolYear) error {
	return r.db.Save(&model).Error
}

func (r *SchoolYearRepoImplementation) FindSchoolYears() (*[]domain.SchoolYear, error) {
	var res []domain.SchoolYear
	if err := r.db.Find(&res).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *SchoolYearRepoImplementation) FindSchoolYearByUuid(uuid string) (*domain.SchoolYear, error) {
	var res domain.SchoolYear
	if err := r.db.First(&res, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *SchoolYearRepoImplementation) DeleteSchoolYear(model *domain.SchoolYear) error {
	return r.db.Delete(&model).Error
}
