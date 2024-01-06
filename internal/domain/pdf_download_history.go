package domain

import "time"

type PdfDownloadHistory struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not_null;varchar(64)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
