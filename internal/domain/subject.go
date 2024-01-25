package domain

import "time"

type Subject struct {
	ID        uint      `gorm:"primaryKey"`
	Uuid      string    `gorm:"not_null;unique"`
	Name      string    `gorm:"not_null;type:varchar(32)"`
	Teachers  []Teacher `gorm:"many2many:teacher_subjects;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
