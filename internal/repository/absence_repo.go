package repository

import "github.com/iki-rumondor/init-golang-service/internal/domain"

type AbsenceRepository interface {
	FindAbsencePagination(pagination *domain.Pagination) (*domain.Pagination, error)
	FindUserByID(uint) (*domain.User, error)
	CheckStudentIsAbsence(studentID, scheduleID uint) int
	FindStudentByUuid(studentUuid string) (*domain.Student, error)
	FindScheduleByUuid(scheduleUuid string) (*domain.Schedule, error)
	FindScheduleByID(uint) (*domain.Schedule, error)
	CreateAbsence(*domain.Absence) error
	FindAbsencesStudent(studentID uint) (*[]domain.Absence, error)

	FindStudentByUserID(userID uint) (*domain.Student, error)
}
