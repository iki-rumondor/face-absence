package repository

import (
	"net/http"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
)

type SchoolFeeRepository interface {
	CreateSchoolFee(model *domain.SchoolFee) error
	FindAllSchoolFees(limit, offset int) (*[]domain.SchoolFee, error)
	FindStudentSchoolFee(studentUuid string) (*[]domain.SchoolFee, error)
	FindBySchoolYear(schoolYearUuid string) (*[]domain.SchoolFee, error)
	FirstStudentSchoolFee(studentUuid string) (*domain.SchoolFee, error)
	FindStudentByUuid(string) (*domain.Student, error)
	FindSchoolYearByUuid(string) (*domain.SchoolYear, error)
	CountStudentSchoolFee(uint, string, string) int
	CheckStudentSchoolFee(studentID, yearID uint, month string) int
	GetSchoolFeesPDF(data *request.SchoolFeePDFData) (*http.Response, error)
	FindSchoolFeeBy(column string, value interface{}) (*domain.SchoolFee, error)
	UpdateSchoolFee(uuid string, req *request.SchoolFee) error
	DeleteSchoolFee(uuid string) error

	GetUtils(key string) (string, error)
	UpdateUtils(key, value string) error
}
