package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/request"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type AbsenceRepoImplementation struct {
	db *gorm.DB
}

func NewAbsenceRepository(db *gorm.DB) AbsenceRepository {
	return &AbsenceRepoImplementation{
		db: db,
	}
}

func (r *AbsenceRepoImplementation) FindAllAbsences() (*[]domain.Absence, error) {
	var absences []domain.Absence
	
	if err := r.db.Preload("Student.Class").Preload("Schedule").Find(&absences).Error; err != nil{
		return nil, err
	}

	return &absences, nil
}

func (r *AbsenceRepoImplementation) FindAbsencePagination(pagination *domain.Pagination) (*domain.Pagination, error) {
	var absence []domain.Absence
	var totalRows int64 = 0

	if err := r.db.Model(&domain.Absence{}).Count(&totalRows).Error; err != nil {
		return nil, err
	}

	if pagination.Limit == 0 {
		pagination.Limit = int(totalRows)
	}

	offset := pagination.Page * pagination.Limit

	if err := r.db.Limit(pagination.Limit).Offset(offset).Preload("Student").Preload("Schedule").Find(&absence).Error; err != nil {
		return nil, err
	}

	var res = []response.AbsenceResponse{}
	for _, item := range absence {
		res = append(res, response.AbsenceResponse{
			Uuid:   item.Uuid,
			Status: item.Status,
			Student: &response.StudentResponse{
				Uuid:         item.Student.Uuid,
				JK:           item.Student.JK,
				NIS:          item.Student.NIS,
				TempatLahir:  item.Student.TempatLahir,
				TanggalLahir: item.Student.TanggalLahir,
				Alamat:       item.Student.Alamat,
				TanggalMasuk: item.Student.TanggalMasuk,
				Image:        item.Student.Image,
			},
			Schedule: &response.ScheduleResponse{
				Uuid:  item.Schedule.Uuid,
				Day:   item.Schedule.Day,
				Start: item.Schedule.Start,
				End:   item.Schedule.End,
			},
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	pagination.Rows = res

	pagination.TotalRows = int(totalRows)

	return pagination, nil
}

func (r *AbsenceRepoImplementation) FindUserByID(id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *AbsenceRepoImplementation) FindScheduleByID(id uint) (*domain.Schedule, error) {
	var schedule domain.Schedule
	if err := r.db.First(&schedule, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &schedule, nil
}

func (r *AbsenceRepoImplementation) CreateAbsence(model *domain.Absence) error {
	return r.db.Create(&model).Error
}

func (r *AbsenceRepoImplementation) FindStudentByUuid(studentUuid string) (*domain.Student, error) {
	var student domain.Student
	if err := r.db.First(&student, "uuid = ?", studentUuid).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *AbsenceRepoImplementation) FindScheduleByUuid(scheduleUuid string) (*domain.Schedule, error) {
	var schedule domain.Schedule
	if err := r.db.First(&schedule, "uuid = ?", scheduleUuid).Error; err != nil {
		return nil, err
	}
	return &schedule, nil
}
func (r *AbsenceRepoImplementation) CheckStudentIsAbsence(studentID, scheduleID uint) int {
	return int(r.db.First(&domain.Absence{}, "student_id = ? AND schedule_id = ?", studentID, scheduleID).RowsAffected)
}

func (r *AbsenceRepoImplementation) FindAbsencesStudent(studentID uint) (*[]domain.Absence, error) {
	var res []domain.Absence
	if err := r.db.Preload("Student").Preload("Schedule.Subject").Find(&res, "student_id", studentID).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *AbsenceRepoImplementation) FindStudentByUserID(userID uint) (*domain.Student, error) {
	var res domain.Student
	if err := r.db.Preload("Class").First(&res, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *AbsenceRepoImplementation) GetAbsencesPDF(data []*request.AbsencePDFData) (*http.Response, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var API_URL = os.Getenv("LARAVEL_API")
	if API_URL == "" {
		return nil, err
	}

	url := fmt.Sprintf("%s/generate-pdf/Daftar_Absensi", API_URL)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	return resp, nil
}
