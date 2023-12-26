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

	if err := r.db.Create(user).Error; err != nil{
		return nil, err
	}

	return user, nil
}

func (r *StudentRepoImplementation) SaveStudent(student *domain.Student) (error) {

	if err := r.db.Save(student).Error; err != nil{
		return err
	}

	return nil
}

func (r *StudentRepoImplementation) DeleteUser(user *domain.User) {

	r.db.Delete(user)
}
