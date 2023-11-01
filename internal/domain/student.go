package domain

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID       uint   `gorm:"primaryKey"`
	NIS      string `gorm:"not_null; unique; varchar(20)"`
	Kelas    string `gorm:"not_null; varchar(10)"`
	JK       string `gorm:"not_null; varchar(10)"`
	Semester string `gorm:"not_null; varchar(5)"`
	User     User
	UserID   uint

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Student) BeforeSave(tx *gorm.DB) error {

	var student Student
	if result := tx.First(&student, "nis = ? AND id != ?", s.NIS, s.ID).RowsAffected; result > 0{
		return errors.New("the nis has already registered")
	}

	return nil
}

type ListOfStudent struct {
	Students []Student
}
