package repository

import (
	"fmt"

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

	if err := r.db.Model(&domain.Subject{}).Count(&totalRows).Error; err != nil {
		return nil, err
	}

	if pagination.Limit == 0 {
		pagination.Limit = int(totalRows)
	}

	offset := pagination.Page * pagination.Limit

	if err := r.db.Limit(pagination.Limit).Offset(offset).Preload("Teachers.User").Find(&subjects).Error; err != nil {
		return nil, err
	}

	var resp []response.SubjectResponse

	for _, res := range subjects {
		var teachers []response.Teacher
		for _, item := range res.Teachers {
			teachers = append(teachers, response.Teacher{
				Uuid:          item.Uuid,
				JK:            item.JK,
				Nip:           item.Nip,
				Nuptk:         item.Nuptk,
				StatusPegawai: item.StatusPegawai,
				TempatLahir:   item.TempatLahir,
				TanggalLahir:  item.TanggalLahir,
				NoHp:          item.NoHp,
				Jabatan:       item.Jabatan,
				TotalJtm:      item.TotalJtm,
				Alamat:        item.Alamat,
				User: &response.UserData{
					Nama:      item.User.Nama,
					Username:  item.User.Username,
					Avatar:    item.User.Avatar,
					CreatedAt: item.User.CreatedAt,
					UpdatedAt: item.User.UpdatedAt,
				},
				CreatedAt: item.CreatedAt,
				UpdatedAt: item.UpdatedAt,
			})
		}

		resp = append(resp, response.SubjectResponse{
			Uuid:      res.Uuid,
			Name:      res.Name,
			Teachers:  &teachers,
			CreatedAt: res.CreatedAt,
			UpdatedAt: res.UpdatedAt,
		})
	}

	pagination.Rows = resp

	pagination.TotalRows = int(totalRows)

	return pagination, nil
}

func (r *SubjectRepoImplementation) CreateSubject(subject *domain.Subject) error {
	return r.db.Create(subject).Error
}

func (r *SubjectRepoImplementation) UpdateSubject(model *domain.Subject) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.db.Model(model).Updates(model).Error; err != nil {
			return err
		}

		if err := r.db.Model(model).Association("Teachers").Replace(model.Teachers); err != nil {
			return err
		}

		return nil
	})
}

func (r *SubjectRepoImplementation) FindSubjects() (*[]domain.Subject, error) {
	var res []domain.Subject
	if err := r.db.Preload("Teachers.User").Find(&res).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *SubjectRepoImplementation) FindSubjectByUuid(uuid string) (*domain.Subject, error) {
	var res domain.Subject
	if err := r.db.Preload("Teachers.User").First(&res, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *SubjectRepoImplementation) DeleteSubject(model *domain.Subject) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(model).Association("Teachers").Clear(); err != nil {
			return err
		}

		if err := r.db.Delete(&model, "uuid = ?", model.Uuid).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *SubjectRepoImplementation) FindTeacherByUuid(uuid string) (*domain.Teacher, error) {
	var res domain.Teacher
	if err := r.db.First(&res, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *SubjectRepoImplementation) FindTeacherBy(column string, value interface{}) (*domain.Teacher, error) {
	var res domain.Teacher
	if err := r.db.First(&res, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &res, nil
}
