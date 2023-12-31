package domain

import (
	"time"
)

type Student struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Uuid         string    `json:"uuid" gorm:"not_null; unique;"`
	NIS          string    `json:"nis" gorm:"not_null; unique; varchar(20)"`
	JK           string    `json:"jk" gorm:"not_null; varchar(10)"`
	TempatLahir  string    `json:"tempat_lahir" gorm:"not_null; varchar(120)"`
	TanggalLahir string    `json:"tanggal_lahir" gorm:"not_null; varchar(120)"`
	Alamat       string    `json:"alamat" gorm:"not_null; varchar(120)"`
	UserID       uint      `json:"user_id"`
	ClassID      uint      `json:"class_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Class        *Class
	User         *User
}
