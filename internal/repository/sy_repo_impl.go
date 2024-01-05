package repository

import (
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
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

func (r *SchoolYearRepoImplementation) FindSchoolYearPagination(pagination *domain.Pagination) (*domain.Pagination, error) {
	var schoolYears []domain.SchoolYear
	var totalRows int64 = 0

	offset := pagination.Page * pagination.Limit

	if err := r.db.Limit(pagination.Limit).Offset(offset).Find(&schoolYears).Error; err != nil {
		return nil, err
	}

	var res = []response.SchoolYearResponse{}
	for _, sy := range schoolYears {
		res = append(res, response.SchoolYearResponse{
			Uuid:      sy.Uuid,
			Name:      sy.Name,
			CreatedAt: sy.CreatedAt,
			UpdatedAt: sy.UpdatedAt,
		})
	}

	pagination.Rows = res

	if err := r.db.Model(&domain.SchoolYear{}).Count(&totalRows).Error; err != nil {
		return nil, err
	}

	pagination.TotalRows = int(totalRows)

	return pagination, nil
}

func (r *SchoolYearRepoImplementation) CreateSchoolYear(model *domain.SchoolYear) error {
	return r.db.Create(model).Error
}

func (r *SchoolYearRepoImplementation) UpdateSchoolYear(model *domain.SchoolYear) error {
	return r.db.Model(model).Where("uuid = ?", model.Uuid).Updates(model).Error
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
	return r.db.Delete(&model, "uuid = ?", model.Uuid).Error
}
