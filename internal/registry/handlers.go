package registry

import (
	customHTTP "github.com/iki-rumondor/init-golang-service/internal/adapter/http"
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
	SchoolFeeHandler  *customHTTP.SchoolFeeHandler
}

func RegisterHandlers(s *Services) *Handlers {

	auth_handler := customHTTP.NewAuthHandler(s.Auth)
	student_handler := customHTTP.NewStudentHandler(s.Student, s.Class)
	teacher_handler := customHTTP.NewTeacherHandler(s.Teacher)
	class_handler := customHTTP.NewClassHandler(s.Class, s.Teacher)
	subject_handler := customHTTP.NewSubjectHandler(s.Subject, s.Teacher)
	sy_handler := customHTTP.NewSchoolYearHandler(s.SchoolYear)
	schedule_handler := customHTTP.NewScheduleHandler(s.Schedule, s.Class, s.Subject, s.Teacher, s.SchoolYear)
	user_handler := customHTTP.NewUserHandler(s.User)
	absence_handler := customHTTP.NewAbsenceHandler(s.Absence, s.Schedule, s.Student)
	school_fee_handler := customHTTP.NewSchoolFeeHandler(s.SchoolFee)

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
		SchoolFeeHandler:  school_fee_handler,
	}
}
