package domain

import (
	"time"

	"github.com/iki-rumondor/init-golang-service/internal/utils"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Nama     string `gorm:"not null; varchar(120)"`
	Username string `gorm:"not null; unique; varchar(120)"`
	Password string `gorm:"not null; varchar(120)"`
	Avatar   *string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeSave(tx *gorm.DB) error {

	if u.Password != "" {
		hashPass, err := utils.HashPassword(u.Password)
		if err != nil {
			return err
		}
		u.Password = hashPass
	}
	return nil
}

func (u *User) BeforeCreate(tx *gorm.DB) error {

	if u.Avatar == nil {
		defaultName := "default-avatar.jpg"
		u.Avatar = &defaultName
	}

	return nil
}
