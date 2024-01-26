package repository

import (
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
)

type ClassRepository interface {
	FindClassPagination(pagination *domain.Pagination) (*domain.Pagination, error)

	CreateClass(*domain.Class) error
	FindClasses() (*[]domain.Class, error)
	FindClassByUuid(string) (*domain.Class, error)
	UpdateClass(*domain.Class) error
	DeleteClass(*domain.Class) error
	
	FindTeacherClassesByUserID(userID uint) (*domain.Teacher, error)

	GetClassPDF(data []*request.ClassPDFData) ([]byte, error)
}
