package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type SubjectRepository interface {
	FindSubjectPagination(pagination *domain.Pagination) (*domain.Pagination, error) 
	FindSubjects() (*[]domain.Subject, error)
	FindSubjectByUuid(string) (*domain.Subject, error)
	
	CreateSubject(*domain.Subject) error
	UpdateSubject(*domain.Subject) error
	DeleteSubject(string) error

	FindTeacherByUuid(uuid string) (*domain.Teacher, error)
	FindTeacherBy(string, interface{}) (*domain.Teacher, error)
}
