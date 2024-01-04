package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type TeacherRepository interface {
	FindTeachersPagination(*domain.Pagination) (*domain.Pagination, error)

	CreateTeacherUser(*domain.Teacher, *domain.User) error
	FindTeachers() (*[]domain.Teacher, error)
	FindTeacherByUuid(string) (*domain.Teacher, error)
	FindTeacherByColumn(column, data string) (*domain.Teacher, error)
	UpdateTeacherUser(*domain.Teacher, *domain.User) error
	DeleteTeacherUser(userID uint) error

	FindUserByUsername(string) (*domain.User, error)
}
