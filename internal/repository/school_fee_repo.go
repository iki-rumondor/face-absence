package repository

import (
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
)

type SchoolFeeRepository interface {
	CreateSchoolFee(req *request.SchoolFee) error
	FindAllSchoolFees(limit, offset int) (*[]domain.SchoolFee, error)
	FindSchoolFeeBy(column string, value interface{}) (*domain.SchoolFee, error)
	UpdateSchoolFee(uuid string, req *request.SchoolFee) error
	DeleteSchoolFee(uuid string) error
}