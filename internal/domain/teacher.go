package domain

import (
	"time"
)

type Teacher struct {
	ID            uint   `gorm:"primaryKey"`
	Uuid          string `gorm:"not_null; unique"`
	Nip           string `gorm:"not_null; unique; varchar(20)"`
	Nuptk         string `gorm:"not_null; varchar(120)"`
	StatusPegawai string `gorm:"not_null; varchar(120)"`
	JK            string `gorm:"not_null; varchar(10)"`
	TempatLahir   string `gorm:"not_null; varchar(120)"`
	TanggalLahir  string `gorm:"not_null; varchar(120)"`
	NoHp          string `gorm:"not_null; varchar(120)"`
	Jabatan       string `gorm:"not_null; varchar(120)"`
	TotalJtm      string `gorm:"not_null; varchar(120)"`
	Alamat        string `gorm:"not_null; varchar(120)"`
	UserID        uint
	User          *User
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
