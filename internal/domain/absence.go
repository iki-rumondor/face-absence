package domain

import (
	"fmt"
	"time"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"gorm.io/gorm"
)

type Absence struct {
	ID         uint   `gorm:"primaryKey"`
	Uuid       string `gorm:"not_null;unique"`
	Status     string `gorm:"not_null; varchar(16)"`
	StudentID  uint
	ScheduleID uint
	Student    *Student
	Schedule   *Schedule
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (m *Absence) BeforeSave(tx *gorm.DB) error {

	if err := tx.First(&Student{}, "id = ?", m.StudentID).Error; err != nil {
		return &response.Error{
			Code:    404,
			Message: fmt.Sprintf("Student with id %d is not found", m.StudentID),
		}
	}

	if err := tx.First(&Schedule{}, "id = ?", m.ScheduleID).Error; err != nil {
		return &response.Error{
			Code:    404,
			Message: fmt.Sprintf("Schedule with id %d is not found", m.ScheduleID),
		}
	}

	return nil
}
