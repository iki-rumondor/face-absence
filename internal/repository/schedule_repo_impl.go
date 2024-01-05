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

	offset := pagination.Page * pagination.Limit

	if err := r.db.Limit(pagination.Limit).Offset(offset).Find(&schedules).Error; err != nil {
		return nil, err
	}

	var res = []response.ScheduleResponse{}
	for _, schedule := range schedules {
		res = append(res, response.ScheduleResponse{
			Uuid:         schedule.Uuid,
			Name:         schedule.Name,
			Day:          schedule.Day,
			Start:        schedule.Start,
			End:          schedule.End,
			ClassID:      schedule.ClassID,
			SubjectID:    schedule.SubjectID,
			TeacherID:    schedule.TeacherID,
			SchoolYearID: schedule.SchoolYearID,
			CreatedAt:    schedule.CreatedAt,
			UpdatedAt:    schedule.UpdatedAt,
		})
	}

	pagination.Rows = res

	if err := r.db.Model(&domain.Schedule{}).Count(&totalRows).Error; err != nil {
		return nil, err
	}

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
	if err := r.db.Find(&res).Error; err != nil {
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
