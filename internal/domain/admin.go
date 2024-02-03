package domain

import "time"

type Admin struct {
	ID        uint   `gorm:"primaryKey"`
	Uuid      string `gorm:"not_null;unique"`
	JK        string `gorm:"not_null;size:16`
	UserID    uint
	User      *User
	CreatedAt time.Time
	UpdatedAt time.Time
}
