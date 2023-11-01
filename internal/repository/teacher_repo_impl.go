package repository

import (
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type TeacherRepoImplementation struct {
	db *gorm.DB
}

func NewTeacherRepository(db *gorm.DB) TeacherRepository {
	return &TeacherRepoImplementation{
		db: db,
	}
}

func (r *TeacherRepoImplementation) CreateUser(user *domain.User) (*domain.User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *TeacherRepoImplementation) CreateTeacher(teacher *domain.Teacher) (*domain.Teacher, error) {
	if err := r.db.Save(teacher).Error; err != nil {
		return nil, err
	}

	return teacher, nil
}

func (r *TeacherRepoImplementation) FindTeachers() (*[]domain.Teacher, error) {
	var teachers []domain.Teacher
	if err := r.db.Preload("User.Role").Find(&teachers).Error; err != nil {
		return nil, err
	}

	return &teachers, nil
}



func (r *TeacherRepoImplementation) DeleteUser(id uint) error {
	if err := r.db.Delete(&domain.User{}, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}
