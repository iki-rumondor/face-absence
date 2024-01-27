package registry

import (
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"gorm.io/gorm"
)

type Repositories struct {
	Absence    repository.AbsenceRepository
	Auth       repository.AuthRepository
	Class      repository.ClassRepository
	Schedule   repository.ScheduleRepository
	Student    repository.StudentRepository
	Subject    repository.SubjectRepository
	SchoolYear repository.SchoolYearRepository
	Teacher    repository.TeacherRepository
	User       repository.UserRepository
	SchoolFee  repository.SchoolFeeRepository
}

func RegisterRepositories(gormDB *gorm.DB) *Repositories {
	auth_repo := repository.NewAuthRepository(gormDB)
	student_repo := repository.NewStudentRepository(gormDB)
	teacher_repo := repository.NewTeacherRepository(gormDB)
	class_repo := repository.NewClassRepository(gormDB)
	subject_repo := repository.NewSubjectRepository(gormDB)
	sy_repo := repository.NewSchoolYearRepository(gormDB)
	schedule_repo := repository.NewScheduleRepository(gormDB)
	user_repo := repository.NewUserRepository(gormDB)
	absence_repo := repository.NewAbsenceRepository(gormDB)
	school_fee_repo := repository.NewSchoolFeeRepository(gormDB)

	return &Repositories{
		Student:    student_repo,
		Auth:       auth_repo,
		Teacher:    teacher_repo,
		Class:      class_repo,
		Subject:    subject_repo,
		SchoolYear: sy_repo,
		Schedule:   schedule_repo,
		User:       user_repo,
		Absence:    absence_repo,
		SchoolFee:  school_fee_repo,
	}

}
