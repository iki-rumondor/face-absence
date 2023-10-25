package repository

import (
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type AuthRepoImplementation struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepoImplementation{
		db: db,
	}
}

func (r *AuthRepoImplementation) FindByEmail(email string) (*domain.User, error){
	var user domain.User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil{
		return nil, err
	}

	return &user, nil
}