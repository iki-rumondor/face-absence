package application

import (
	"errors"
	"log"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
	"gorm.io/gorm"
)

type SchoolFeeService struct {
	Repo repository.SchoolFeeRepository
}

func NewSchoolFeeService(repo repository.SchoolFeeRepository) *SchoolFeeService {
	return &SchoolFeeService{
		Repo: repo,
	}
}

func (s *SchoolFeeService) CreateSchoolFee(req *request.SchoolFee) error {

	if err := s.Repo.CreateSchoolFee(req); err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &response.Error{
				Code:    404,
				Message: "Siswa dengan uuid yang dimasukkan tidak ditemukan",
			}
		}
		return INTERNAL_ERROR
	}

	return nil
}

func (s *SchoolFeeService) GetAllSchoolFees(url string, page, limit int) (*[]domain.SchoolFee, map[string]interface{}, error) {
	offset := page * limit

	schoolFees, err := s.Repo.FindAllSchoolFees(limit, offset)
	if err != nil {
		log.Println(err.Error())
		return nil, nil, INTERNAL_ERROR
	}

	pagination, err := utils.CalculatePagination(url, len(*schoolFees), page, limit)
	if err != nil {
		log.Println(err.Error())
		return nil, nil, INTERNAL_ERROR
	}

	return schoolFees, pagination, nil
}

func (s *SchoolFeeService) GetSchoolFeeByUuid(uuid string) (*domain.SchoolFee, error) {

	schoolFee, err := s.Repo.FindSchoolFeeBy("uuid", uuid)
	if err != nil {
		log.Println(err.Error())
		return nil, INTERNAL_ERROR
	}

	return schoolFee, nil
}

func (s *SchoolFeeService) UpdateSchoolFee(uuid string, req *request.SchoolFee) error {

	if err := s.Repo.UpdateSchoolFee(uuid, req); err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &response.Error{
				Code:    404,
				Message: "Siswa dengan uuid yang dimasukkan tidak ditemukan",
			}
		}
		return INTERNAL_ERROR
	}

	return nil
}

func (s *SchoolFeeService) DeleteSchoolFee(uuid string) error {

	if err := s.Repo.DeleteSchoolFee(uuid); err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &response.Error{
				Code:    404,
				Message: "Siswa dengan uuid yang dimasukkan tidak ditemukan",
			}
		}
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return &response.Error{
				Code:    404,
				Message: "Gagal Menghapus: Data ini berelasi dengan data yang lain",
			}
		}
		return INTERNAL_ERROR
	}

	return nil
}
