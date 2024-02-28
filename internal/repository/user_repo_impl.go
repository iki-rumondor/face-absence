package repository

import (
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type UserRepoImplementation struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepoImplementation{
		db: db,
	}
}

func (r *UserRepoImplementation) FindUserByID(ID uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, "id = ?", ID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepoImplementation) FindTeacherByUuid(uuid string) (*domain.Teacher, error) {
	var teacher domain.Teacher
	if err := r.db.Preload("User").First(&teacher, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &teacher, nil
}

func (r *UserRepoImplementation) UpdateAvatar(model *domain.User) error {
	return r.db.Model(model).Where("id = ?", model.ID).Update("avatar", model.Avatar).Error
}

func (r *UserRepoImplementation) Update(model *domain.User) error {
	return r.db.Updates(model).Error
}

func (r *UserRepoImplementation) FindUsers() (*[]domain.User, error) {
	var user []domain.User
	if err := r.db.Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepoImplementation) CountStudentsTeachersAdmins() (map[string]int64, error) {
	var (
		count_admin       int64
		count_student     int64
		count_teacher     int64
		count_admin_man   int64
		count_student_man int64
		count_teacher_man int64
	)

	if err := r.db.Model(&domain.Admin{}).Count(&count_admin).Error; err != nil {
		return nil, err
	}

	if err := r.db.Model(&domain.Student{}).Count(&count_student).Error; err != nil {
		return nil, err
	}

	if err := r.db.Model(&domain.Teacher{}).Count(&count_teacher).Error; err != nil {
		return nil, err
	}

	if err := r.db.Find(&domain.Admin{}, "jk = ?", "LAKI-LAKI").Count(&count_admin_man).Error; err != nil {
		return nil, err
	}

	if err := r.db.Find(&domain.Student{}, "jk = ?", "LAKI-LAKI").Count(&count_student_man).Error; err != nil {
		return nil, err
	}

	if err := r.db.Find(&domain.Teacher{}, "jk = ?", "LAKI-LAKI").Count(&count_teacher_man).Error; err != nil {
		return nil, err
	}

	resp := map[string]int64{
		"admin":       count_admin,
		"student":     count_student,
		"teacher":     count_teacher,
		"admin_man":   count_admin_man,
		"student_man": count_student_man,
		"teacher_man": count_teacher_man,
	}

	return resp, nil
}
