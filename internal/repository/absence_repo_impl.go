package repository

import (
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
