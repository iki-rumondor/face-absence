package domain

import (
	"time"
)

type Schedule struct {
	ID           uint   `gorm:"primaryKey"`
	Uuid         string `gorm:"not_null;unique"`
	Day          string `gorm:"not_null;varchar(16)"`
	Start        string `gorm:"not_null;varchar(5)"`
	End          string `gorm:"not_null;varchar(5)"`
	ClassID      uint
	SubjectID    uint
	SchoolYearID uint
	Class        *Class
	Subject      *Subject
	SchoolYear   *SchoolYear
	Absences     *[]Absence
	// Students     []Student `gorm:"many2many:student_schedules;"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
