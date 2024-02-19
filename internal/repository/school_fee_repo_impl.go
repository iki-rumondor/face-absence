package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
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

func (r *SchoolFeeRepoImplementation) CreateSchoolFee(model *domain.SchoolFee) error {
	return r.db.Create(&model).Error
}

func (r *SchoolFeeRepoImplementation) FindAllSchoolFees(limit, offset int) (*[]domain.SchoolFee, error) {
	var schoolFees []domain.SchoolFee

	query := r.db.Offset(offset).Preload("Student.Class").Preload("SchoolYear")
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
	if err := r.db.Preload("Student.Class").Preload("SchoolYear").First(&schoolFee, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}

	return &schoolFee, nil
}

func (r *SchoolFeeRepoImplementation) FindStudentSchoolFee(studentUuid string) (*[]domain.SchoolFee, error) {
	var student domain.Student
	if err := r.db.First(&student, "uuid = ?", studentUuid).Error; err != nil {
		return nil, err
	}

	var schoolFee []domain.SchoolFee
	if err := r.db.Preload("Student.Class").Preload("SchoolYear").Find(&schoolFee, "student_id = ?", student.ID).Error; err != nil {
		return nil, err
	}

	return &schoolFee, nil
}

func (r *SchoolFeeRepoImplementation) FirstStudentSchoolFee(studentUuid string) (*domain.SchoolFee, error) {
	var student domain.Student
	if err := r.db.First(&student, "uuid = ?", studentUuid).Error; err != nil {
		return nil, err
	}

	var schoolFee domain.SchoolFee
	if err := r.db.Preload("Student.Class").Preload("SchoolYear").Order("created_at desc").First(&schoolFee, "student_id = ?", student.ID).Error; err != nil {
		return nil, err
	}

	return &schoolFee, nil
}

func (r *SchoolFeeRepoImplementation) FindBySchoolYear(schoolYearUuid string) (*[]domain.SchoolFee, error) {
	var sy domain.SchoolYear
	if err := r.db.First(&sy, "uuid = ?", schoolYearUuid).Error; err != nil {
		return nil, err
	}

	var schoolFee []domain.SchoolFee
	if err := r.db.Preload("Student.Class").Preload("SchoolYear").Find(&schoolFee, "school_year_id = ?", sy.ID).Error; err != nil {
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

	var sy domain.SchoolYear
	if err := r.db.First(&sy, "uuid = ?", req.SchoolYearUuid).Error; err != nil {
		return err
	}

	date, err := utils.FormatToTime(req.Date, "2006-01-02")
	if err != nil {
		return err
	}

	model := domain.SchoolFee{
		ID:           schoolFee.ID,
		Date:         date,
		Month:        req.Month,
		Status:       req.Status,
		StudentID:    student.ID,
		SchoolYearID: sy.ID,
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

func (r *SchoolFeeRepoImplementation) GetSchoolFeesPDF(data *request.SchoolFeePDFData) (*http.Response, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var API_URL = os.Getenv("LARAVEL_API")
	if API_URL == "" {
		return nil, err
	}

	url := fmt.Sprintf("%s/school-fee-pdf", API_URL)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *SchoolFeeRepoImplementation) FindStudentByUuid(uuid string) (*domain.Student, error) {
	var result domain.Student
	if err := r.db.First(&result, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *SchoolFeeRepoImplementation) FindSchoolYearByUuid(uuid string) (*domain.SchoolYear, error) {
	var result domain.SchoolYear
	if err := r.db.First(&result, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *SchoolFeeRepoImplementation) CountStudentSchoolFee(studentID uint, month, year string) int {
	return int(r.db.First(&domain.SchoolFee{}, "student_id = ? AND YEAR(date) = ? AND MONTH(date) = ?", studentID, month, year).RowsAffected)
}

func (r *SchoolFeeRepoImplementation) GetUtils(key string) (string, error) {
	var result domain.Utils
	if err := r.db.First(&result, "name = ?", key).Error; err != nil {
		return "", err
	}

	return result.Value, nil
}

func (r *SchoolFeeRepoImplementation) UpdateUtils(name, value string) error {
	return r.db.Model(&domain.Utils{}).Where("name = ?", name).Update("value", value).Error
}
