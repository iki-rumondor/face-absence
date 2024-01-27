package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SchoolFee struct {
	ID        uint   `gorm:"primaryKey"`
	Uuid      string `gorm:"not_null; unique;"`
	Date      string `gorm:"not_null; varchar(20)"`
	Nominal   int    `gorm:"not_null; varchar(10)"`
	StudentID uint
	Student   *Student
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *SchoolFee) BeforeCreate(tx *gorm.DB) error {

	s.Uuid = uuid.NewString()
	return nil
}
