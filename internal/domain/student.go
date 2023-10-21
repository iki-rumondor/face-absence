package domain

import "time"

type Student struct {
	ID    uint   `gorm:"primaryKey"`
	Uuid  string `gorm:"not null; unique; varchar(120)"`
	Nama  string `gorm:"not_null; varchar(120)"`
	NIS   string `gorm:"not_null; varchar(20)"`
	Kelas string `gorm:"not_null; varchar(10)"`
	// JK uint	`gorm:"not_null"`
	// Semester uint `gorm:"not_null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type ListOfStudent struct{
	Students []Student
}
