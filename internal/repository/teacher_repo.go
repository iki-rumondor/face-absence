package repository

import (
	"net/http"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
)

type TeacherRepository interface {
	FindTeachersPagination(*domain.Pagination) (*domain.Pagination, error)

	CreateTeacherUser(*domain.Teacher, *domain.User) error
	FindTeachers() (*[]domain.Teacher, error)
	FindTeacherByUuid(string) (*domain.Teacher, error)
	FindTeacherByColumn(column, data string) (*domain.Teacher, error)
	UpdateTeacherUser(*domain.Teacher, *domain.User) error
	DeleteTeacher(string) error
	GetTeachersPDF(data []*request.TeacherPDFData) (*http.Response, error)
	FindUserByUsername(string) (*domain.User, error)
}
