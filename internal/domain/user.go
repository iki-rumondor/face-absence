package domain

import "time"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Uuid     string `gorm:"not null; unique; varchar(120)"`
	Nama     string `gorm:"not null; varchar(120)"`
	Email    string `gorm:"not null; unique; varchar(120)"`
	Password string `gorm:"not null; varchar(120)"`
	RoleID   uint
	Student  Student
	Role Role `gorm:"foreignKey:RoleID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Role struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null; varchar(120)"`
}

type ListOfUsers struct {
	Users []User
}
