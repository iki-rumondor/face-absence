package domain

import (
	"fmt"
	"time"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"gorm.io/gorm"
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

func (m *Student) BeforeSave(tx *gorm.DB) error {

	if err := tx.First(&Class{}, "id = ?", m.ClassID).Error; err != nil {
		return &response.Error{
			Code:    404,
			Message: fmt.Sprintf("Teacher with id %d is not found", m.ClassID),
		}
	}

	return nil
}
