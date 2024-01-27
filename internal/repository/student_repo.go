package repository

import (
	"net/http"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
)

type StudentRepository interface {
	CreateStudent(student *domain.Student) error
	PaginationStudents(pagination *domain.Pagination) (*domain.Pagination, error)
	FindAllStudents() (*[]domain.Student, error)
	FindStudent(string) (*domain.Student, error)
	FindStudentByUserID(uint) (*domain.Student, error)
	UpdateStudent(student *domain.Student) error
	UpdateStudentImage(uuid, imagePath string) error
	DeleteStudent(student *domain.Student) error

	FindClassBy(column string, value interface{}) (*domain.Class, error)

	GetStudentsPDF(data []*request.StudentPDFData) (*http.Response, error)
	CreateBatchStudents(*[]domain.Student, string) error

	FindLatestHistory() (*domain.PdfDownloadHistory, error)
	CreatePdfHistory(*domain.PdfDownloadHistory) error
}
