package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SchoolFee struct {
	ID           uint      `gorm:"primaryKey"`
	Uuid         string    `gorm:"not_null;unique;"`
	Date         time.Time `gorm:"not_null;type:date"`
	Nominal      int       `gorm:"not_null;varchar(10);default:1150000"`
	Month        string    `gorm:"not_null"`
	Status       string    `gorm:"not_null"`
	SchoolYearID uint      `gorm:"not_null"`
	StudentID    uint
	Student      *Student
	SchoolYear   *SchoolYear
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (s *SchoolFee) BeforeCreate(tx *gorm.DB) error {
	s.Uuid = uuid.NewString()
	return nil
}
