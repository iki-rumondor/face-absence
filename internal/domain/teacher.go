package domain

import (
	"time"
)

type Teacher struct {
	ID            uint    `gorm:"primaryKey"`
	Uuid          string  `gorm:"not_null; unique"`
	Nip           *string `gorm:"unique;varchar(20)"`
	Nuptk         *string `gorm:"varchar(120)"`
	StatusPegawai string  `gorm:"not_null; varchar(120)"`
	JK            string  `gorm:"not_null; varchar(10)"`
	TempatLahir   string  `gorm:"not_null; varchar(120)"`
	TanggalLahir  string  `gorm:"not_null; varchar(120)"`
	NoHp          string  `gorm:"not_null; varchar(120)"`
	Jabatan       string  `gorm:"not_null; varchar(120)"`
	TotalJtm      string  `gorm:"not_null; varchar(120)"`
	Alamat        string  `gorm:"not_null; varchar(120)"`
	UserID        uint
	User          *User
	Classes       *[]Class
	Subjects      []Subject `gorm:"many2many:teacher_subjects;"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// func (m *Teacher) BeforeUpdate(tx *gorm.DB) error {

// 	if m.Nip != nil && *m.Nip == "" {
// 		if err := tx.Exec("UPDATE teachers SET nip = NULL WHERE uuid = ?", m.Uuid).Error; err != nil {
// 			return err
// 		}
// 	}

// 	if m.Nuptk != nil && *m.Nuptk == "" {
// 		if err := tx.Exec("UPDATE teachers SET nuptk = NULL WHERE uuid = ?", m.Uuid).Error; err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
