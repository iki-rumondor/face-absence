package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type TeacherRepository interface {
	CreateUser(*domain.User) (*domain.User, error)
	CreateTeacher(*domain.Teacher) (*domain.Teacher, error)

	DeleteUser(uint) error
}