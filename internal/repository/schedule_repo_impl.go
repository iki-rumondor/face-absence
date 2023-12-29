package repository

import (
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

func (r *ScheduleRepoImplementation) CreateSchedule(model *domain.Schedule) error {
	return r.db.Create(model).Error
}

func (r *ScheduleRepoImplementation) UpdateSchedule(model *domain.Schedule) error {
	return r.db.Model(model).Where("id = ?", model.ID).Updates(model).Error
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
	return r.db.Delete(&model).Error
}
