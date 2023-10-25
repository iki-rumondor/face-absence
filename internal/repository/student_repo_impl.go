package repository

import (
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

func (r *StudentRepoImplementation) SaveList(students *domain.ListOfStudent) error {
	return r.db.Save(students.Students).Error
}

func (r *StudentRepoImplementation) FindAll() (*domain.ListOfStudent, error) {
	var students domain.ListOfStudent
	if err := r.db.Find(&students.Students).Error; err != nil {
		return nil, err
	}
	return &students, nil
}

func (r *StudentRepoImplementation) Find(uuid string) (*domain.Student, error) {
	var student domain.Student

	if err := r.db.First(&student, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &student, nil
}

func (r *StudentRepoImplementation) Save(student *domain.Student) error {
	
	if err := r.db.Save(&student).Error; err != nil {
		return err
	}

	return nil
}

func (r *StudentRepoImplementation) Delete(student *domain.Student) error {
	
	if err := r.db.Delete(student).Error; err != nil {
		return err
	}

	return nil
}
