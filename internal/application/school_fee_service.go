package application

import (
	"errors"
	"fmt"
	"io"
	"log"
	"strings"

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

	student, err := s.Repo.FindStudentByUuid(req.StudentUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &response.Error{
				Code:    404,
				Message: "Siswa dengan uuid yang dimasukkan tidak ditemukan",
			}
		}
		return INTERNAL_ERROR
	}

	dateParts := strings.Split(req.Date, "-")

	if result := s.Repo.CountStudentSchoolFee(student.ID, dateParts[0], dateParts[1]); result > 0 {
		return &response.Error{
			Code:    400,
			Message: "Santri Sudah Melakukan Pembayaran SPP Pada Bulan Ini",
		}
	}

	date, err := utils.FormatToTime(req.Date, "2006-01-02")
	if err != nil {
		return err
	}

	schoolYear, err := s.Repo.FindSchoolYearByUuid(req.SchoolYearUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &response.Error{
				Code:    404,
				Message: "Tahun Pelajaran dengan uuid yang dimasukkan tidak ditemukan",
			}
		}
		return INTERNAL_ERROR
	}

	model := domain.SchoolFee{
		Date:         date,
		// Nominal:      req.Nominal,
		Month:        req.Month,
		SchoolYearID: schoolYear.ID,
		StudentID:    student.ID,
	}

	if err := s.Repo.CreateSchoolFee(&model); err != nil {
		log.Println(err.Error())
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

func (s *SchoolFeeService) GetStudentSchoolFee(studentUuid string) (*[]domain.SchoolFee, error) {

	schoolFees, err := s.Repo.FindStudentSchoolFee(studentUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.Error{
				Code:    404,
				Message: fmt.Sprintf("SPP Santri dengan uuid %s tidak ditemukan", studentUuid),
			}
		}
		log.Println(err.Error())
		return nil, INTERNAL_ERROR
	}

	return schoolFees, nil
}

func (s *SchoolFeeService) GetBySchoolYear(studentUuid string) (*[]domain.SchoolFee, error) {

	schoolFees, err := s.Repo.FindBySchoolYear(studentUuid)
	if err != nil {
		log.Println(err.Error())
		return nil, INTERNAL_ERROR
	}

	return schoolFees, nil
}

func (s *SchoolFeeService) GetNewStudentSchoolFee(studentUuid string) (*domain.SchoolFee, error) {

	schoolFees, err := s.Repo.FirstStudentSchoolFee(studentUuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &response.Error{
				Code:    404,
				Message: fmt.Sprintf("SPP Santri dengan uuid %s tidak ditemukan", studentUuid),
			}
		}
		log.Println(err.Error())
		return nil, INTERNAL_ERROR
	}

	return schoolFees, nil
}

func (s *SchoolFeeService) UpdateSchoolFee(uuid string, req *request.SchoolFee) error {

	if err := s.Repo.UpdateSchoolFee(uuid, req); err != nil {
		log.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &response.Error{
				Code:    404,
				Message: "Data Tidak Ditemukan",
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
				Message: "Data tidak ditemukan",
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

func (s *SchoolFeeService) CreateSchoolFeesPDF(studentUuid string) ([]byte, error) {
	schoolFees, err := s.Repo.FindStudentSchoolFee(studentUuid)
	if err != nil {
		return nil, err
	}

	if len(*schoolFees) == 0 {
		return nil, &response.Error{
			Code:    404,
			Message: "Data SPP Masih Kosong",
		}
	}

	var fee []request.SchoolFeeData
	var name, class string

	for _, item := range *schoolFees {
		name = item.Student.Nama
		class = item.Student.Class.Name
		fee = append(fee, request.SchoolFeeData{
			Nominal: item.Nominal,
			Date:    item.Date.Format("02-01-2006"),
			Month:   utils.GetBulanIndonesia(item.Date.Format("01")),
		})
	}

	data := request.SchoolFeePDFData{
		StudentName:   name,
		Class:         class,
		SchoolFeeData: fee,
	}

	resp, err := s.Repo.GetSchoolFeesPDF(&data)
	if err != nil {
		log.Println(err.Error())
		return nil, INTERNAL_ERROR
	}

	defer resp.Body.Close()

	pdfData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, INTERNAL_ERROR
	}

	return pdfData, nil
}
