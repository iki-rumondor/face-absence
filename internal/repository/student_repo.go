package repository

import (
	"github.com/iki-rumondor/init-golang-service/internal/domain"
)

type StudentRepository interface {
	// SaveList(*domain.ListOfStudent) error
	SaveStudent(*domain.Student) error
	FindAllStudentUsers() (*[]domain.Student, error)
	FindStudentUser(string) (*domain.Student, error)
	FindStudent(string) (*domain.Student, error)
	UpdateStudentUser(*domain.Student, *domain.User) error
	DeleteStudentUser(*domain.Student, *domain.User) error

	CreateUser(*domain.User) (*domain.User, error)
	DeleteUser(*domain.User) error
}
