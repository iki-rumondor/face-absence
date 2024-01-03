package domain

import "time"

type Admin struct {
	ID        uint   `gorm:"primaryKey"`
	Uuid      string `gorm:"not_null;unique"`
	UserID    uint
	User      *User
	CreatedAt time.Time
	UpdatedAt time.Time
}
