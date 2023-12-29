package repository

import (
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type SubjectRepoImplementation struct {
	db *gorm.DB
}

func NewSubjectRepository(db *gorm.DB) SubjectRepository {
	return &SubjectRepoImplementation{
		db: db,
	}
}

func (r *SubjectRepoImplementation) CreateSubject(model *domain.Subject) error {
	return r.db.Create(model).Error
}

func (r *SubjectRepoImplementation) UpdateSubject(model *domain.Subject) error {
	return r.db.Model(model).Where("id = ?", model.ID).Updates(model).Error
}

func (r *SubjectRepoImplementation) FindSubjects() (*[]domain.Subject, error) {
	var res []domain.Subject
	if err := r.db.Find(&res).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *SubjectRepoImplementation) FindSubjectByUuid(uuid string) (*domain.Subject, error) {
	var res domain.Subject
	if err := r.db.First(&res, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *SubjectRepoImplementation) DeleteSubject(model *domain.Subject) error {
	return r.db.Delete(&model).Error
}
