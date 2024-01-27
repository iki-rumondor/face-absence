package repository

import (
	"fmt"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type SchoolFeeRepoImplementation struct {
	db *gorm.DB
}

func NewSchoolFeeRepository(db *gorm.DB) SchoolFeeRepository {
	return &SchoolFeeRepoImplementation{
		db: db,
	}
}

func (r *SchoolFeeRepoImplementation) CreateSchoolFee(req *request.SchoolFee) error {
	var student domain.Student
	if err := r.db.First(&student, "uuid = ?", req.StudentUuid).Error; err != nil {
		return err
	}

	schoolFee := domain.SchoolFee{
		Date:      req.Date,
		Nominal:   req.Nominal,
		StudentID: student.ID,
	}

	return r.db.Create(&schoolFee).Error
}

func (r *SchoolFeeRepoImplementation) FindAllSchoolFees(limit, offset int) (*[]domain.SchoolFee, error) {
	var schoolFees []domain.SchoolFee

	query := r.db.Offset(offset).Preload("Student.Class")
	if limit > 0 {
		query = query.Limit(limit)
	}

	if err := query.Find(&schoolFees).Error; err != nil {
		return nil, err
	}

	return &schoolFees, nil
}

func (r *SchoolFeeRepoImplementation) FindSchoolFeeBy(column string, value interface{}) (*domain.SchoolFee, error) {
	var schoolFee domain.SchoolFee
	if err := r.db.Preload("Student.Class").First(&schoolFee, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &schoolFee, nil
}

func (r *SchoolFeeRepoImplementation) UpdateSchoolFee(uuid string, req *request.SchoolFee) error {
	var schoolFee domain.SchoolFee
	if err := r.db.First(&schoolFee, "uuid = ?", uuid).Error; err != nil {
		return err
	}

	var student domain.Student
	if err := r.db.First(&student, "uuid = ?", req.StudentUuid).Error; err != nil {
		return err
	}

	model := domain.SchoolFee{
		ID:        schoolFee.ID,
		Date:      req.Date,
		Nominal:   req.Nominal,
		StudentID: student.ID,
	}

	return r.db.Model(&model).Updates(&model).Error
}

func (r *SchoolFeeRepoImplementation) DeleteSchoolFee(uuid string) error {
	var schoolFee domain.SchoolFee
	if err := r.db.First(&schoolFee, "uuid = ?", uuid).Error; err != nil {
		return err
	}

	return r.db.Delete(&schoolFee).Error
}
