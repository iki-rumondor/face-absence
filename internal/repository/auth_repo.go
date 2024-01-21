package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type AuthRepository interface {
	FindByEmail(string) (*domain.User, error)
	FindByUsername(string) (*domain.User, error)
	FindTeacherByUserID(uint) error
	FindStudentByUserID(uint) error
	FindAdminByUserID(uint) error
	FindUserByID(id uint) (*domain.User, error)
}
