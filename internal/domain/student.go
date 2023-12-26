package domain

import (
	"time"
)

type Student struct {
	ID           uint    `gorm:"primaryKey"`
	Uuid         string  `gorm:"not_null; unique;"`
	NIS          string  `gorm:"not_null; unique; varchar(20)"`
	JK           string  `gorm:"not_null; varchar(10)"`
	TempatLahir  string  `gorm:"not_null; varchar(120)"`
	TanggalLahir string  `gorm:"not_null; varchar(120)"`
	Alamat       string  `gorm:"not_null; varchar(120)"`
	UserID       uint
	ClassID      uint
	Class        *Class
	User         *User
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
