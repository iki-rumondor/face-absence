package repository

import (
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
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

func (r *SubjectRepoImplementation) FindSubjectPagination(pagination *domain.Pagination) (*domain.Pagination, error) {
	var subjects []domain.Subject
	var totalRows int64 = 0

	offset := pagination.Page * pagination.Limit

	if err := r.db.Limit(pagination.Limit).Offset(offset).Find(&subjects).Error; err != nil {
		return nil, err
	}

	var res = []response.SubjectResponse{}
	for _, subject := range subjects {
		res = append(res, response.SubjectResponse{
			Uuid:          subject.Uuid,
			Name:          subject.Name,
			CreatedAt:     subject.CreatedAt,
			UpdatedAt:     subject.UpdatedAt,
		})
	}

	pagination.Rows = res

	if err := r.db.Model(&domain.Subject{}).Count(&totalRows).Error; err != nil {
		return nil, err
	}

	pagination.TotalRows = int(totalRows)

	return pagination, nil
}

func (r *SubjectRepoImplementation) CreateSubject(model *domain.Subject) error {
	return r.db.Create(model).Error
}

func (r *SubjectRepoImplementation) UpdateSubject(model *domain.Subject) error {
	return r.db.Model(model).Where("uuid = ?", model.Uuid).Updates(model).Error
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
	return r.db.Delete(&model, "uuid = ?", model.Uuid).Error
}
