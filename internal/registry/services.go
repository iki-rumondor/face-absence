package registry

import "github.com/iki-rumondor/init-golang-service/internal/application"

type Services struct {
	Absence    *application.AbsenceService
	Auth       *application.AuthService
	Class      *application.ClassService
	Schedule   *application.ScheduleService
	Student    *application.StudentService
	Subject    *application.SubjectService
	SchoolYear *application.SchoolYearService
	Teacher    *application.TeacherService
	User       *application.UserService
}

func RegisterServices(repo *Repositories) *Services {
	auth_service := application.NewAuthService(repo.Auth)
	student_service := application.NewStudentService(repo.Student)
	teacher_service := application.NewTeacherService(repo.Teacher)
	class_service := application.NewClassService(repo.Class)
	subject_service := application.NewSubjectService(repo.Subject)
	sy_service := application.NewSchoolYearService(repo.SchoolYear)
	schedule_service := application.NewScheduleService(repo.Schedule)
	user_service := application.NewUserService(repo.User)
	absence_service := application.NewAbsenceService(repo.Absence)

	return &Services{
		Student:    student_service,
		Auth:       auth_service,
		Teacher:    teacher_service,
		Class:      class_service,
		Subject:    subject_service,
		SchoolYear: sy_service,
		Schedule:   schedule_service,
		User:       user_service,
		Absence:    absence_service,
	}

}
