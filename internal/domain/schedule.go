package domain

import (
	"fmt"
	"time"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"gorm.io/gorm"
)

type Schedule struct {
	ID           uint   `gorm:"primaryKey"`
	Uuid         string `gorm:"not_null;unique"`
	Name         string `gorm:"not_null;varchar(32)"`
	Day          string `gorm:"not_null;varchar(10)"`
	Start        string `gorm:"not_null;varchar(5)"`
	End          string `gorm:"not_null;varchar(5)"`
	ClassID      uint
	SubjectID    uint
	TeacherID    uint
	SchoolYearID uint
	Class        *Class
	Subject      *Subject
	Teacher      *Teacher
	SchoolYear   *SchoolYear
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (m *Schedule) BeforeSave(tx *gorm.DB) error {

	if err := tx.First(&Class{}, "id = ?", m.ClassID).Error; err != nil {
		return &response.Error{
			Code:    404,
			Message: fmt.Sprintf("Class with id %d is not found", m.ClassID),
		}
	}

	if err := tx.First(&Subject{}, "id = ?", m.SubjectID).Error; err != nil {
		return &response.Error{
			Code:    404,
			Message: fmt.Sprintf("Subject with id %d is not found", m.SubjectID),
		}
	}

	if err := tx.First(&Teacher{}, "id = ?", m.TeacherID).Error; err != nil {
		return &response.Error{
			Code:    404,
			Message: fmt.Sprintf("Teacher with id %d is not found", m.TeacherID),
		}
	}

	if err := tx.First(&SchoolYear{}, "id = ?", m.SchoolYearID).Error; err != nil {
		return &response.Error{
			Code:    404,
			Message: fmt.Sprintf("School Year with id %d is not found", m.SchoolYearID),
		}
	}

	return nil
}
