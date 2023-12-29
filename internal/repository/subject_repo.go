package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type SubjectRepository interface {
	FindSubjects() (*[]domain.Subject, error)
	FindSubjectByUuid(string) (*domain.Subject, error)
	SaveSubject(*domain.Subject) error
	DeleteSubject(*domain.Subject) error
}
