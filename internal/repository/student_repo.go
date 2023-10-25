package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type StudentRepository interface{
	SaveList(*domain.ListOfStudent) error
	Save(*domain.Student) error
	FindAll() (*domain.ListOfStudent, error)
	Find(string) (*domain.Student, error)
	Delete(*domain.Student) error
}
