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
	if err := r.db.First(&user, "id = ?", ID).Error; err != nil{
		return nil, err
	}
	return &user, nil
}

func (r *UserRepoImplementation) UpdateAvatar(model *domain.User) error {
	return r.db.Model(model).Where("id = ?", model.ID).Update("avatar", model.Avatar).Error
}
