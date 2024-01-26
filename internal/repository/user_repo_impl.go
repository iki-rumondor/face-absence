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

func (r *UserRepoImplementation) UpdateAvatar(model *domain.User) error {
	return r.db.Model(model).Where("id = ?", model.ID).Update("avatar", model.Avatar).Error
}
func (r *UserRepoImplementation) CountStudentsTeachersAdmins() (map[string]int64, error) {
	var (
		count_admin   int64
		count_student int64
		count_teacher int64
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

	resp := map[string]int64{
		"admin": count_admin,
		"student": count_student,
		"teacher": count_teacher,
	}

	return resp, nil
}
