package repository

import (
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"gorm.io/gorm"
)

type StudentRepoImplementation struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository{
	return &StudentRepoImplementation{
		db: db,
	}
}

func (r *StudentRepoImplementation) Save(student []*domain.Student) error{
	return r.db.Save(student).Error
}