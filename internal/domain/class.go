package domain

import (
	"time"
)

type Class struct {
	ID        uint   `gorm:"primaryKey"`
	Uuid      string `gorm:"not_null;unique"`
	Name      string `gorm:"not_null;varchar(32)"`
	TeacherID uint
	Teacher   *Teacher
	CreatedAt time.Time
	UpdatedAt time.Time
}
