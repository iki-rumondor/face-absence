package domain

import (
	"errors"
	"time"

	"github.com/iki-rumondor/init-golang-service/internal/utils"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Uuid     string `gorm:"not null; unique; varchar(120)"`
	Nama     string `gorm:"not null; varchar(120)"`
	Email    string `gorm:"not null; unique; varchar(120)"`
	Password string `gorm:"not null; varchar(120)"`
	RoleID   uint
	Role     Role `gorm:"foreignKey:RoleID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeSave(tx *gorm.DB) error {

	var user User
	if result := tx.First(&user, "email = ? AND id != ?", u.Email, u.ID).RowsAffected; result > 0{
		return errors.New("the email has already been taken")
	}

	hashPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashPass
	return nil
}

type ListOfUsers struct {
	Users []User
}

type Role struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null; varchar(120)"`
}
