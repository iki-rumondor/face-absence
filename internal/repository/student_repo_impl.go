package repository

import (
	"fmt"
	"math"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type StudentRepoImplementation struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &StudentRepoImplementation{
		db: db,
	}
}

func (r *StudentRepoImplementation) CreateStudent(student *domain.Student) error {
	return r.db.Create(student).Error
}

func (r *StudentRepoImplementation) PaginationStudents(pagination *domain.Pagination) (*domain.Pagination, error) {
	var students []domain.Student

	var totalPages, fromRow, toRow = 0, 0, 0
	var totalRows int64 = 0

	if err := r.db.Model(&domain.Student{}).Count(&totalRows).Error; err != nil {
		return nil, err
	}

	if pagination.Limit == 0 {
		pagination.Limit = int(totalRows)
	}

	offset := pagination.Page * pagination.Limit

	if err := r.db.Limit(pagination.Limit).Offset(offset).Preload("User").Preload("Class").Find(&students).Error; err != nil {
		return nil, err
	}

	var res = []response.StudentResponse{}
	for _, student := range students {
		res = append(res, response.StudentResponse{
			Uuid:         student.Uuid,
			JK:           student.JK,
			NIS:          student.NIS,
			TempatLahir:  student.TempatLahir,
			TanggalLahir: student.TanggalLahir,
			Alamat:       student.Alamat,
			Class: &response.ClassData{
				Uuid:      student.Class.Uuid,
				Name:      student.Class.Name,
				CreatedAt: student.Class.CreatedAt,
				UpdatedAt: student.Class.UpdatedAt,
			},
			CreatedAt: student.CreatedAt,
			UpdatedAt: student.UpdatedAt,
		})
	}

	pagination.Rows = res

	pagination.TotalRows = int(totalRows)

	totalPages = int(math.Ceil(float64(totalRows)/float64(pagination.Limit)) - 1)
	pagination.TotalPages = totalPages

	if pagination.Page == 0 {
		fromRow = 1
		toRow = pagination.Limit
	} else {
		if pagination.Page <= totalPages {
			fromRow = pagination.Page*pagination.Limit + 1
			toRow = pagination.Page + 1*pagination.Limit
		}
	}

	if toRow > int(totalRows) {
		toRow = int(totalRows)
	}

	pagination.FromRow = fromRow
	pagination.ToRow = toRow

	return pagination, nil
}

func (r *StudentRepoImplementation) FindAllStudents() (*[]domain.Student, error) {
	var students []domain.Student
	if err := r.db.Preload("User").Preload("Class").Find(&students).Error; err != nil {
		return nil, err
	}
	return &students, nil
}

func (r *StudentRepoImplementation) FindStudent(uuid string) (*domain.Student, error) {
	var student domain.Student
	if err := r.db.Preload("User").Preload("Class").First(&student, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepoImplementation) FindStudentByUserID(userID uint) (*domain.Student, error) {
	var student domain.Student
	if err := r.db.Preload("User").Preload("Class").First(&student, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepoImplementation) UpdateStudent(student *domain.Student) error {
	return r.db.Model(student).Updates(student).Error
}

func (r *StudentRepoImplementation) DeleteStudent(student *domain.Student) error {
	return r.db.Delete(student).Error
}

func (r *StudentRepoImplementation) FindLatestHistory() (*domain.PdfDownloadHistory, error) {
	var history domain.PdfDownloadHistory
	if err := r.db.Last(&history).Error; err != nil {
		return nil, err
	}
	return &history, nil
}

func (r *StudentRepoImplementation) FindClassBy(column string, value interface{}) (*domain.Class, error) {
	var class domain.Class
	if err := r.db.First(&class, fmt.Sprintf("%s = ?", column), value).Error; err != nil {
		return nil, err
	}
	return &class, nil
}

func (r *StudentRepoImplementation) CreatePdfHistory(history *domain.PdfDownloadHistory) error {
	return r.db.Create(history).Error
}
