package repository

import (
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

func (r *ClassRepoImplementation) SaveClass(class *domain.Class) error {
	return r.db.Save(&class).Error
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
	return r.db.Delete(&class).Error
}