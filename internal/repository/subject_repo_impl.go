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

	if err := r.db.Limit(pagination.Limit).Offset(offset).Preload("Teacher.User").Find(&subjects).Error; err != nil {
		return nil, err
	}

	var res = []response.SubjectResponse{}
	for _, subject := range subjects {
		res = append(res, response.SubjectResponse{
			Uuid: subject.Uuid,
			Name: subject.Name,
			// Teacher: &response.Teacher{
			// 	Uuid:          subject.Teacher.Uuid,
			// 	JK:            subject.Teacher.JK,
			// 	Nip:           subject.Teacher.Nip,
			// 	Nuptk:         subject.Teacher.Nuptk,
			// 	StatusPegawai: subject.Teacher.StatusPegawai,
			// 	TempatLahir:   subject.Teacher.TempatLahir,
			// 	TanggalLahir:  subject.Teacher.TanggalLahir,
			// 	NoHp:          subject.Teacher.NoHp,
			// 	Jabatan:       subject.Teacher.Jabatan,
			// 	TotalJtm:      subject.Teacher.TotalJtm,
			// 	Alamat:        subject.Teacher.Alamat,
			// 	User: &response.UserData{
			// 		Nama:      subject.Teacher.User.Nama,
			// 		Username:  subject.Teacher.User.Username,
			// 		Avatar:    subject.Teacher.User.Avatar,
			// 		CreatedAt: subject.Teacher.User.CreatedAt,
			// 		UpdatedAt: subject.Teacher.User.UpdatedAt,
			// 	},
			// 	CreatedAt: subject.Teacher.CreatedAt,
			// 	UpdatedAt: subject.Teacher.UpdatedAt,
			// },
			CreatedAt: subject.CreatedAt,
			UpdatedAt: subject.UpdatedAt,
		})
	}

	pagination.Rows = res

	pagination.TotalRows = int(totalRows)

	return pagination, nil
}

func (r *SubjectRepoImplementation) CreateSubject(subject *domain.Subject) error {
	return r.db.Create(subject).Error
}

func (r *SubjectRepoImplementation) UpdateSubject(model *domain.Subject) error {
	return r.db.Model(model).Updates(model).Error
}

func (r *SubjectRepoImplementation) FindSubjects() (*[]domain.Subject, error) {
	var res []domain.Subject
	if err := r.db.Preload("Teacher.User").Find(&res).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *SubjectRepoImplementation) FindSubjectByUuid(uuid string) (*domain.Subject, error) {
	var res domain.Subject
	if err := r.db.Preload("Teacher.User").First(&res, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *SubjectRepoImplementation) DeleteSubject(model *domain.Subject) error {
	return r.db.Delete(&model, "uuid = ?", model.Uuid).Error
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
