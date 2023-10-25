package domain

import "time"

type Student struct {
	ID       uint   `gorm:"primaryKey"`
	NIS      string `gorm:"not_null; varchar(20)"`
	Kelas    string `gorm:"not_null; varchar(10)"`
	JK       string `gorm:"not_null; varchar(10)"`
	Semester string `gorm:"not_null; varchar(5)"`
	UserID   uint

	CreatedAt time.Time
	UpdatedAt time.Time
}

type ListOfStudent struct {
	Students []Student
}
