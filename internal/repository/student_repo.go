package repository

import (
	"github.com/iki-rumondor/init-golang-service/internal/domain"
)

type StudentRepository interface {
	CreateStudentUser(*domain.Student, *domain.User) error
	PaginationStudents(pagination *domain.Pagination) (*domain.Pagination, error)
	FindAllStudents() (*[]domain.Student, error)
	FindStudent(string) (*domain.Student, error)
	UpdateStudent(*domain.Student, *domain.User) error
	DeleteStudent(userID uint) error

	CreateUser(*domain.User) (*domain.User, error)
	SaveStudent(*domain.Student) error
	DeleteUser(*domain.User)

	
}
