package domain

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Teacher struct {
	ID     uint   `gorm:"primaryKey"`
	Nip    string `gorm:"not_null; unique; varchar(20)"`
	JK     string `gorm:"not_null; varchar(10)"`
	UserID uint
	User   User

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Teacher) BeforeSave(tx *gorm.DB) error {

	var teacher Teacher
	if result := tx.First(&teacher, "nip = ? AND id != ?", t.Nip, t.ID).RowsAffected; result > 0 {
		return errors.New("the nip has already registered")
	}

	return nil
}
