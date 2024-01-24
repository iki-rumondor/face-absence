package domain

import (
	"time"
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
	SchoolYearID uint
	Class        *Class
	Subject      *Subject
	SchoolYear   *SchoolYear
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
