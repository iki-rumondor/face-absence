package repository

import (
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type ClassRepoImplementation struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) ClassRepository {
	return &ClassRepoImplementation{
		db: db,
	}
}

func (r *ClassRepoImplementation) FindClassPagination(pagination *domain.Pagination) (*domain.Pagination, error) {
	var classes []domain.Class
	var totalRows int64 = 0

	if err := r.db.Model(&domain.Class{}).Count(&totalRows).Error; err != nil {
		return nil, err
	}

	if pagination.Limit == 0 {
		pagination.Limit = int(totalRows)
	}
	
	offset := pagination.Page * pagination.Limit

	if err := r.db.Limit(pagination.Limit).Offset(offset).Preload("Teacher").Find(&classes).Error; err != nil {
		return nil, err
	}

	var res = []response.ClassResponse{}
	for _, class := range classes {
		res = append(res, response.ClassResponse{
			Uuid:      class.Uuid,
			Name:      class.Name,
			TeacherID: class.TeacherID,
			CreatedAt: class.CreatedAt,
			UpdatedAt: class.UpdatedAt,
		})
	}

	pagination.Rows = res

	pagination.TotalRows = int(totalRows)

	return pagination, nil
}

func (r *ClassRepoImplementation) CreateClass(class *domain.Class) error {
	return r.db.Create(class).Error
}

func (r *ClassRepoImplementation) UpdateClass(class *domain.Class) error {
	return r.db.Model(class).Where("uuid = ?", class.Uuid).Updates(class).Error
}

func (r *ClassRepoImplementation) FindClasses() (*[]domain.Class, error) {
	var classes []domain.Class
	if err := r.db.Find(&classes).Error; err != nil {
		return nil, err
	}

	return &classes, nil
}

func (r *ClassRepoImplementation) FindClassByUuid(uuid string) (*domain.Class, error) {
	var class domain.Class
	if err := r.db.First(&class, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &class, nil
}

func (r *ClassRepoImplementation) DeleteClass(class *domain.Class) error {
	return r.db.Delete(&class, "uuid = ?", class.Uuid).Error
}
