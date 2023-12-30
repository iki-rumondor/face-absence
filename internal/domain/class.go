package domain

import (
	"fmt"
	"time"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"
	"gorm.io/gorm"
)

type Class struct {
	ID        uint   `gorm:"primaryKey"`
	Uuid      string `gorm:"not_null;unique"`
	Name      string `gorm:"not_null;varchar(32)"`
	TeacherID uint
	Teacher   *Teacher
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Class) BeforeSave(tx *gorm.DB) error {

	if err := tx.First(&Teacher{}, "id = ?", m.TeacherID).Error; err != nil {
		return &response.Error{
			Code: 404,
			Message: fmt.Sprintf("Teacher with id %d is not found", m.TeacherID),
		}
	}

	return nil
}
