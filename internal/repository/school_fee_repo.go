package repository

import (
	"net/http"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
)

type SchoolFeeRepository interface {
	CreateSchoolFee(req *request.SchoolFee) error
	FindAllSchoolFees(limit, offset int) (*[]domain.SchoolFee, error)
	FindStudentSchoolFee(studentUuid string) (*[]domain.SchoolFee, error)
	GetSchoolFeesPDF(data *request.SchoolFeePDFData) (*http.Response, error)
	FindSchoolFeeBy(column string, value interface{}) (*domain.SchoolFee, error)
	UpdateSchoolFee(uuid string, req *request.SchoolFee) error
	DeleteSchoolFee(uuid string) error
}
