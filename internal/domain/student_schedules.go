package domain

import "time"

type StudentSchedules struct {
	ID         uint   `gorm:"primaryKey"`
	Uuid       string `gorm:"not_null;unique"`
	StudentID  uint
	ScheduleID uint
	Student    *Student
	Schedule   *Schedule
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
