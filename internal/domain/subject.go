package domain

import "time"

type Subject struct {
	ID        uint   `gorm:"primaryKey"`
	Uuid      string `gorm:"not_null;unique"`
	Name      string `gorm:"not_null;varchar(32)"`
	TeacherID uint
	Teacher   *Teacher
	Schedules  *[]Schedule
	CreatedAt time.Time
	UpdatedAt time.Time
}
