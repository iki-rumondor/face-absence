package routes

import (
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
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "ngrok-skip-browser-warning"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12,
	}))

	router.MaxMultipartMemory = 10 << 20

	public := router.Group("api")
	{
		public.POST("/auth/login", handlers.AuthHandler.Login)
		public.GET("/auth/verify-token", middleware.IsValidJWT(), handlers.AuthHandler.VerifyToken)
	}

	student := router.Group("api").Use(middleware.IsValidJWT(), middleware.IsStudent())
	{
		student.PATCH("users/avatar", middleware.SetUserID(), handlers.UserHandler.UpdateAvatar)
		student.POST("absence", middleware.SetUserID(), handlers.AbsenceHandler.CreateAbsence)
	}

	admin := router.Group("api").Use(middleware.IsValidJWT(), middleware.IsAdmin())
	{
		admin.POST("master/teachers", handlers.TeacherHandler.CreateTeacher)
		admin.GET("master/teachers", handlers.TeacherHandler.GetTeachersPagination)
		admin.GET("master/teachers/:uuid", handlers.TeacherHandler.GetTeacher)
		admin.PUT("master/teachers/:uuid", handlers.TeacherHandler.UpdateTeacher)
		admin.DELETE("master/teachers/:uuid", handlers.TeacherHandler.DeleteTeacher)

		admin.POST("master/students", handlers.StudentHandler.ImportStudentsData)
		admin.GET("master/students", handlers.StudentHandler.GetAllStudentsData)
		admin.GET("master/students/:uuid", handlers.StudentHandler.GetStudentData)
		admin.PUT("master/students/:uuid", handlers.StudentHandler.UpdateStudentData)
		admin.DELETE("master/students/:uuid", handlers.StudentHandler.DeleteStudent)

		admin.POST("master/classes", handlers.ClassHandler.CreateClass)
		admin.GET("master/classes", handlers.ClassHandler.GetClassPagination)
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
	}

	return router
}
