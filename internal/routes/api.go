package routes

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/middleware"
	"github.com/iki-rumondor/init-golang-service/internal/registry"
)

func StartServer(handlers *registry.Handlers) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "ngrok-skip-browser-warning", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12,
	}))

	router.MaxMultipartMemory = 10 << 20

	router.StaticFS("/public/avatar", http.Dir("internal/assets/avatar"))
	router.StaticFS("/public/file", http.Dir("internal/assets/temp"))

	public := router.Group("api")
	{
		public.POST("/auth/login", handlers.AuthHandler.Login)
		public.GET("/auth/verify-token", middleware.IsValidJWT(), middleware.SetUserID(), middleware.SetUserRole(), handlers.AuthHandler.VerifyToken)
	}

	teacher := router.Group("api").Use(middleware.IsValidJWT(), middleware.IsRole("GURU"))
	{
		teacher.GET("teacher/classes", middleware.SetUserID(), handlers.ClassHandler.GetTeacherClasses)
		teacher.GET("teacher/classes/:uuid", middleware.SetUserID(), handlers.ClassHandler.GetClassWithStudents)
		teacher.PATCH("users/avatar", middleware.SetUserID(), handlers.UserHandler.UpdateAvatar)
		// teacher.POST("absence", middleware.SetUserID(), handlers.AbsenceHandler.CreateAbsence)
		teacher.GET("schedules", middleware.SetUserID(), handlers.ScheduleHandler.GetTeacherSchedules)
		teacher.GET("schedules/:uuid", handlers.ScheduleHandler.GetScheduleForStudent)
		teacher.GET("absences/history", middleware.SetUserID(), handlers.AbsenceHandler.GetStudentAbsences)
	}

	user := router.Group("api").Use(middleware.IsValidJWT())
	{
		// user.POST("master/students/import", handlers.StudentHandler.ImportStudentsData)
		// user.GET("master/students/report", handlers.StudentHandler.CreateReport)
		user.GET("dashboard", handlers.UserHandler.GetDashboardData)
		user.PATCH("students/:uuid/image", handlers.StudentHandler.UpdateStudentImage)

		user.POST("master/students/import", handlers.StudentHandler.ImportStudentsData)
		user.POST("master/students", handlers.StudentHandler.CreateStudent)
		user.GET("master/students", handlers.StudentHandler.GetAllStudentsData)
		user.GET("master/students/:uuid", handlers.StudentHandler.GetStudentData)
		user.PUT("master/students/:uuid", handlers.StudentHandler.UpdateStudentData)
		user.DELETE("master/students/:uuid", handlers.StudentHandler.DeleteStudent)
	}

	admin := router.Group("api").Use(middleware.IsValidJWT(), middleware.IsRole("ADMIN"))
	{
		admin.POST("master/teachers", handlers.TeacherHandler.CreateTeacher)
		admin.GET("master/teachers", handlers.TeacherHandler.GetTeachersPagination)
		admin.GET("master/teachers/:uuid", handlers.TeacherHandler.GetTeacher)
		admin.PUT("master/teachers/:uuid", handlers.TeacherHandler.UpdateTeacher)
		admin.DELETE("master/teachers/:uuid", handlers.TeacherHandler.DeleteTeacher)

		admin.POST("master/classes", handlers.ClassHandler.CreateClass)
		admin.GET("master/classes", handlers.ClassHandler.GetClassPagination)
		admin.GET("master/classes/option", handlers.ClassHandler.GetClassOption)
		admin.GET("master/classes/:uuid", handlers.ClassHandler.GetClass)
		admin.PUT("master/classes/:uuid", handlers.ClassHandler.UpdateClass)
		admin.DELETE("master/classes/:uuid", handlers.ClassHandler.DeleteClass)

		admin.POST("master/subjects", handlers.SubjectHandler.CreateSubject)
		admin.GET("master/subjects", handlers.SubjectHandler.GetSubjectPagination)
		admin.GET("master/subjects/:uuid", handlers.SubjectHandler.GetSubject)
		admin.PUT("master/subjects/:uuid", handlers.SubjectHandler.UpdateSubject)
		admin.DELETE("master/subjects/:uuid", handlers.SubjectHandler.DeleteSubject)

		admin.POST("master/school_years", handlers.SchoolYearHandler.CreateSchoolYear)
		admin.GET("master/school_years", handlers.SchoolYearHandler.GetSchoolYearPagination)
		admin.GET("master/school_years/:uuid", handlers.SchoolYearHandler.GetSchoolYear)
		admin.PUT("master/school_years/:uuid", handlers.SchoolYearHandler.UpdateSchoolYear)
		admin.DELETE("master/school_years/:uuid", handlers.SchoolYearHandler.DeleteSchoolYear)

		admin.POST("master/schedules", handlers.ScheduleHandler.CreateSchedule)
		admin.GET("master/schedules", handlers.ScheduleHandler.GetSchedulePagination)
		admin.GET("master/schedules/:uuid", handlers.ScheduleHandler.GetSchedule)
		admin.PUT("master/schedules/:uuid", handlers.ScheduleHandler.UpdateSchedule)
		admin.DELETE("master/schedules/:uuid", handlers.ScheduleHandler.DeleteSchedule)

		admin.GET("absences", handlers.AbsenceHandler.GetAllAbsences)

		admin.GET("pdf/classes", handlers.ClassHandler.GetClassPDF)
		admin.GET("pdf/students", handlers.StudentHandler.GetStudentsPDF)
		admin.GET("pdf/teachers", handlers.TeacherHandler.GetTeachersPDF)
		// admin.GET("download/:filename", customHTTP.DownloadFile)
	}

	return router
}
