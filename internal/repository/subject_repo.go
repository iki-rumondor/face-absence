package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type SubjectRepository interface {
	FindSubjectPagination(pagination *domain.Pagination) (*domain.Pagination, error) 
	CreateSubject(*domain.Subject) error
	FindSubjects() (*[]domain.Subject, error)
	FindSubjectByUuid(string) (*domain.Subject, error)
	UpdateSubject(*domain.Subject) error
	DeleteSubject(*domain.Subject) error
}
