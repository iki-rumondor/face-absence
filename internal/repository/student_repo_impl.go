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

func (r *StudentRepoImplementation) FindAllStudentUsers() (*[]domain.User, error) {
	var users []domain.User
	if err := r.db.Preload("Student").Preload("Role").Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *StudentRepoImplementation) FindStudentUser(uuid string) (*domain.User, error) {
	var user domain.User

	if err := r.db.Preload("Student").Preload("Role").First(&user, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *StudentRepoImplementation) FindStudent(uuid string) (*domain.Student, error) {
	var user domain.User
	var student domain.Student

	if err := r.db.First(&user, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	if err := r.db.First(&student, "user_id = ?", user.ID).Error; err != nil {
		return nil, err
	}

	return &student, nil
}

func (r *StudentRepoImplementation) UpdateStudentUser(student *domain.Student, user *domain.User) error {

	err := r.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Model(student).Updates(student).Error; err != nil {
			return err
		}

		if err := tx.Model(user).Updates(user).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *StudentRepoImplementation) DeleteStudentUser(student *domain.Student, user *domain.User) error {

	err := r.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Delete(&student).Error; err != nil {
			return err
		}

		if err := tx.Delete(&user).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
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
