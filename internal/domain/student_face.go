package domain

import (
	"time"
)

type StudentFace struct {
	ID         uint   `gorm:"primaryKey"`
	FaceEncode string `gorm:"not_null"`
	StudentID  uint
	Student    *Student
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
