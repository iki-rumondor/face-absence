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

	if err := r.db.Limit(pagination.Limit).Offset(offset).Preload("Class").Preload("Subject").Preload("SchoolYear").Find(&schedules).Error; err != nil {
		return nil, err
	}

	var res = []response.ScheduleResponse{}
	for _, schedule := range schedules {
		res = append(res, response.ScheduleResponse{
			Uuid:  schedule.Uuid,
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
			SchoolYear: &response.SchoolYearResponse{
				Uuid:      schedule.SchoolYear.Uuid,
				Name:      schedule.SchoolYear.Name,
				CreatedAt: schedule.SchoolYear.CreatedAt,
				UpdatedAt: schedule.SchoolYear.UpdatedAt,
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
	if err := r.db.Preload("Class").Preload("Subject").Preload("SchoolYear").Find(&res).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *ScheduleRepoImplementation) FindSchedulesByClass(classID uint) (*[]domain.Schedule, error) {
	var res []domain.Schedule
	if err := r.db.Preload("Class").Preload("Subject").Preload("SchoolYear").Find(&res, "class_id", classID).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *ScheduleRepoImplementation) FindScheduleByUuid(uuid string) (*domain.Schedule, error) {
	var res domain.Schedule
	if err := r.db.Preload("Class").Preload("Subject").Preload("SchoolYear").First(&res, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *ScheduleRepoImplementation) DeleteSchedule(model *domain.Schedule) error {
	return r.db.Delete(&model, "uuid = ?", model.Uuid).Error
}

func (r *ScheduleRepoImplementation) FindStudentByUserID(userID uint) (*domain.Student, error) {
	var res domain.Student
	if err := r.db.Preload("Class").First(&res, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *ScheduleRepoImplementation) FindUserByID(ID uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, "id = ?", ID).Error; err != nil{
		return nil, err
	}
	return &user, nil
}

func (r *ScheduleRepoImplementation) FindStudentAbsenceByScheduleID(studentID, scheduleID uint) (*domain.Absence, error) {
	var absence domain.Absence
	if err := r.db.Preload("Student").Preload("Schedule").First(&absence, "student_id = ? AND schedule_id = ?", studentID, scheduleID).Error; err != nil{
		return nil, err
	}
	return &absence, nil
}