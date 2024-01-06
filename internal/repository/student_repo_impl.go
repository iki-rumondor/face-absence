package repository

import (
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

func (r *StudentRepoImplementation) CreateStudentUser(student *domain.Student, user *domain.User) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		student.UserID = user.ID
		if err := tx.Create(student).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *StudentRepoImplementation) PaginationStudents(pagination *domain.Pagination) (*domain.Pagination, error) {
	var students []domain.Student

	var totalPages, fromRow, toRow = 0, 0, 0
	var totalRows int64 = 0

	offset := pagination.Page * pagination.Limit

	if err := r.db.Limit(pagination.Limit).Offset(offset).Preload("User").Find(&students).Error; err != nil {
		return nil, err
	}

	var res = []response.StudentUser{}
	for _, student := range students {
		res = append(res, response.StudentUser{
			Uuid:         student.Uuid,
			Nama:         student.User.Nama,
			Username:     student.User.Username,
			JK:           student.JK,
			NIS:          student.NIS,
			TempatLahir:  student.TempatLahir,
			TanggalLahir: student.TanggalLahir,
			Alamat:       student.Alamat,
			UserID:       student.UserID,
			ClassID:      student.ClassID,
			CreatedAt:    student.CreatedAt,
			UpdatedAt:    student.UpdatedAt,
		})
	}

	pagination.Rows = res

	if err := r.db.Model(&domain.Student{}).Count(&totalRows).Error; err != nil {
		return nil, err
	}

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
	if err := r.db.Preload("User").Find(&students).Error; err != nil {
		return nil, err
	}
	return &students, nil
}

func (r *StudentRepoImplementation) FindStudent(uuid string) (*domain.Student, error) {
	var student domain.Student
	if err := r.db.Preload("User").First(&student, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepoImplementation) UpdateStudent(student *domain.Student, user *domain.User) error {

	return r.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Model(student).Updates(student).Error; err != nil {
			return err
		}

		if err := tx.Model(user).Updates(user).Error; err != nil {
			return err
		}

		return nil
	})

}

func (r *StudentRepoImplementation) DeleteStudent(userID uint) error {

	return r.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Delete(&domain.Student{}, "user_id = ?", userID).Error; err != nil {
			return err
		}

		if err := tx.Delete(&domain.User{}, "id = ?", userID).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *StudentRepoImplementation) CreateUser(user *domain.User) (*domain.User, error) {

	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *StudentRepoImplementation) SaveStudent(student *domain.Student) error {

	if err := r.db.Save(student).Error; err != nil {
		return err
	}

	return nil
}

func (r *StudentRepoImplementation) DeleteUser(user *domain.User) {

	r.db.Delete(user)
}
