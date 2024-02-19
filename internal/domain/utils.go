package domain

import (
	"time"
)

type Utils struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Value     string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
