package domain

import (
	"fmt"
	"time"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"gorm.io/gorm"
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

func (m *Student) BeforeSave(tx *gorm.DB) error {

	if err := tx.First(&Class{}, "id = ?", m.ClassID).Error; err != nil {
		return &response.Error{
			Code: 404,
			Message: fmt.Sprintf("Teacher with id %d is not found", m.ClassID),
		}
	}

	return nil
}
