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

func (r *TeacherRepoImplementation) SaveUser(user *domain.User) (*domain.User, error) {
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

func (r *TeacherRepoImplementation) FindTeacher(uuid string) (*domain.Teacher, error) {
	var teacher domain.Teacher
	if err := r.db.Joins("User").Preload("User.Role").First(&teacher, "User.uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &teacher, nil
}

func (r *TeacherRepoImplementation) UpdateTeacher(user *domain.User, teacher *domain.Teacher) (*domain.Teacher, error) {

	var result domain.Teacher

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(user).Where("uuid = ?", user.Uuid).Update("nama", user.Nama).Error; err != nil {
			return err
		}

		if err := tx.Model(teacher).Where("nip = ?", teacher.Nip).Updates(teacher).First(&result).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *TeacherRepoImplementation) FindUserByUuid(uuid string) (*domain.User, error) {
	var user domain.User

	if err := r.db.First(&user, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &user, nil

}

func (r *TeacherRepoImplementation) DeleteWithUser(user *domain.User) error {

	err := r.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Delete(&domain.Teacher{}, "user_id = ?", user.ID).Error; err != nil {
			return err
		}

		if err := tx.Delete(user).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *TeacherRepoImplementation) DeleteUser(id uint) error {
	if err := r.db.Delete(&domain.User{}, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}
