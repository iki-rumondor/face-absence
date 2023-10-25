package domain

import "time"

type User struct {
	ID       uint `gorm:"primaryKey"`
	Uuid     string `gorm:"not null; unique; varchar(120)"`
	Username string `gorm:"not null; unique; varchar(120)"`
	Email    string `gorm:"not null; unique; varchar(120)"`
	Password string `gorm:"not null; unique; varchar(120)"`
	RoleID   uint

	CreatedAt time.Time
	UpdatedAt time.Time
}

