package registry

import (
	customHTTP "github.com/iki-rumondor/init-golang-service/internal/adapter/http"
	"github.com/iki-rumondor/init-golang-service/internal/application"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"gorm.io/gorm"
)

type Handlers struct {
	StudentHandler    *customHTTP.StudentHandlers
	AuthHandler       *customHTTP.AuthHandlers
	TeacherHandler    *customHTTP.TeacherHandlers
	ClassHandler      *customHTTP.ClassHandler
	SubjectHandler    *customHTTP.SubjectHandler
	SchoolYearHandler *customHTTP.SchoolYearHandler
	ScheduleHandler   *customHTTP.ScheduleHandler
	UserHandler       *customHTTP.UserHandler
	AbsenceHandler    *customHTTP.AbsenceHandler
}

func RegisterHandlers(gormDB *gorm.DB) *Handlers {
	auth_repo := repository.NewAuthRepository(gormDB)
	auth_service := application.NewAuthService(auth_repo)
	auth_handler := customHTTP.NewAuthHandler(auth_service)

	student_repo := repository.NewStudentRepository(gormDB)
	student_service := application.NewStudentService(student_repo)
	student_handler := customHTTP.NewStudentHandler(student_service)

	teacher_repo := repository.NewTeacherRepository(gormDB)
	teacher_service := application.NewTeacherService(teacher_repo)
	teacher_handler := customHTTP.NewTeacherHandler(teacher_service)

	class_repo := repository.NewClassRepository(gormDB)
	class_service := application.NewClassService(class_repo)
	class_handler := customHTTP.NewClassHandler(class_service)

	subject_repo := repository.NewSubjectRepository(gormDB)
	subject_service := application.NewSubjectService(subject_repo)
	subject_handler := customHTTP.NewSubjectHandler(subject_service)

	sy_repo := repository.NewSchoolYearRepository(gormDB)
	sy_service := application.NewSchoolYearService(sy_repo)
	sy_handler := customHTTP.NewSchoolYearHandler(sy_service)

	schedule_repo := repository.NewScheduleRepository(gormDB)
	schedule_service := application.NewScheduleService(schedule_repo)
	schedule_handler := customHTTP.NewScheduleHandler(schedule_service)

	user_repo := repository.NewUserRepository(gormDB)
	user_service := application.NewUserService(user_repo)
	user_handler := customHTTP.NewUserHandler(user_service)

	absence_repo := repository.NewAbsenceRepository(gormDB)
	absence_service := application.NewAbsenceService(absence_repo)
	absence_handler := customHTTP.NewAbsenceHandler(absence_service)

	return &Handlers{
		StudentHandler:    student_handler,
		AuthHandler:       auth_handler,
		TeacherHandler:    teacher_handler,
		ClassHandler:      class_handler,
		SubjectHandler:    subject_handler,
		SchoolYearHandler: sy_handler,
		ScheduleHandler:   schedule_handler,
		UserHandler:       user_handler,
		AbsenceHandler:    absence_handler,
	}
}
