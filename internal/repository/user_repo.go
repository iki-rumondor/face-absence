package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type UserRepository interface {
	FindUserByID(uint) (*domain.User, error)
	UpdateAvatar(*domain.User) error
	CountStudentsTeachersAdmins() (map[string]int64, error)
}
