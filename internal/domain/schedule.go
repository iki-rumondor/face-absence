package domain

import (
	"time"
)

type Schedule struct {
	ID           uint   `gorm:"primaryKey"`
	Uuid         string `gorm:"not_null;unique"`
	Name         string `gorm:"not_null;varchar(32)"`
	Day          time.Time
	Start        time.Time
	End          time.Time
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
