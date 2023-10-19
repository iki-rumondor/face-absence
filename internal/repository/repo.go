package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type StudentRepository interface{
	Save([]*domain.Student) error
}
