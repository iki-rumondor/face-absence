package repository

import (
	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type ScheduleRepoImplementation struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &ScheduleRepoImplementation{
		db: db,
	}
}

func (r *ScheduleRepoImplementation) FindSchedulePagination(pagination *domain.Pagination) (*domain.Pagination, error) {
	var schedules []domain.Schedule
	var totalRows int64 = 0

	if err := r.db.Model(&domain.Schedule{}).Count(&totalRows).Error; err != nil {
		return nil, err
	}

	if pagination.Limit == 0 {
		pagination.Limit = int(totalRows)
	}

	offset := pagination.Page * pagination.Limit

	if err := r.db.Limit(pagination.Limit).Offset(offset).Preload("Class").Preload("Subject").Preload("Teacher").Preload("SchoolYear").Find(&schedules).Error; err != nil {
		return nil, err
	}

	var res = []response.ScheduleResponse{}
	for _, schedule := range schedules {
		res = append(res, response.ScheduleResponse{
			Uuid:  schedule.Uuid,
			Name:  schedule.Name,
			Day:   schedule.Day,
			Start: schedule.Start,
			End:   schedule.End,
			Class: &response.ClassData{
				Uuid:      schedule.Class.Uuid,
				Name:      schedule.Class.Name,
				CreatedAt: schedule.Class.CreatedAt,
				UpdatedAt: schedule.Class.UpdatedAt,
			},
			Subject: &response.SubjectResponse{
				Uuid:      schedule.Subject.Uuid,
				Name:      schedule.Subject.Name,
				CreatedAt: schedule.Subject.CreatedAt,
				UpdatedAt: schedule.Subject.UpdatedAt,
			},
			Teacher: &response.TeacherData{
				Uuid:          schedule.Teacher.Uuid,
				JK:            schedule.Teacher.JK,
				Nip:           schedule.Teacher.Nip,
				Nuptk:         schedule.Teacher.Nuptk,
				StatusPegawai: schedule.Teacher.StatusPegawai,
				TempatLahir:   schedule.Teacher.TempatLahir,
				TanggalLahir:  schedule.Teacher.TanggalLahir,
				NoHp:          schedule.Teacher.NoHp,
				Jabatan:       schedule.Teacher.Jabatan,
				TotalJtm:      schedule.Teacher.TotalJtm,
				Alamat:        schedule.Teacher.Alamat,
				CreatedAt:     schedule.Teacher.CreatedAt,
				UpdatedAt:     schedule.Teacher.UpdatedAt,
			},
			SchoolYear: &response.SchoolYearResponse{
				Uuid:      schedule.Subject.Uuid,
				Name:      schedule.Subject.Name,
				CreatedAt: schedule.Subject.CreatedAt,
				UpdatedAt: schedule.Subject.UpdatedAt,
			},
			CreatedAt: schedule.CreatedAt,
			UpdatedAt: schedule.UpdatedAt,
		})
	}

	pagination.Rows = res

	pagination.TotalRows = int(totalRows)

	return pagination, nil
}

func (r *ScheduleRepoImplementation) CreateSchedule(model *domain.Schedule) error {
	return r.db.Create(model).Error
}

func (r *ScheduleRepoImplementation) UpdateSchedule(model *domain.Schedule) error {
	return r.db.Model(model).Where("uuid = ?", model.ID).Updates(model).Error
}

func (r *ScheduleRepoImplementation) FindSchedules() (*[]domain.Schedule, error) {
	var res []domain.Schedule
	if err := r.db.Preload("Class").Preload("Subject").Preload("Teacher").Preload("SchoolYear").Find(&res).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *ScheduleRepoImplementation) FindScheduleByUuid(uuid string) (*domain.Schedule, error) {
	var res domain.Schedule
	if err := r.db.First(&res, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *ScheduleRepoImplementation) DeleteSchedule(model *domain.Schedule) error {
	return r.db.Delete(&model, "uuid = ?", model.Uuid).Error
}
