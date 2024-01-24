package repository

import (
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

func (r *AbsenceRepoImplementation) CheckStudentIsAbsence(studentID, scheduleID uint) int {
	return int(r.db.First(&domain.Absence{}, "student_id = ? AND schedule_id = ?", studentID, scheduleID).RowsAffected)
}
