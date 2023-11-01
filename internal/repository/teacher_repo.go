package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type TeacherRepository interface {
	SaveUser(*domain.User) (*domain.User, error)
	CreateTeacher(*domain.Teacher) (*domain.Teacher, error)
	FindTeachers() (*[]domain.Teacher, error)
	FindTeacher(string) (*domain.Teacher, error)
	FindUserByUuid(string) (*domain.User, error)
	UpdateTeacher(*domain.User, *domain.Teacher) (*domain.Teacher, error)
	DeleteWithUser(*domain.User) error

	DeleteUser(uint) error
}
