package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	customHTTP "github.com/iki-rumondor/init-golang-service/internal/adapter/http"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/middleware"
)

func StartServer(handlers *customHTTP.Handlers) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
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

	admin := router.Group("api").Use(middleware.IsValidJWT(), middleware.IsAdmin())
	{
		admin.GET("master/teachers", handlers.TeacherHandler.GetTeachers)
		admin.GET("master/teacher/:uuid", handlers.TeacherHandler.GetTeacher)
		admin.POST("master/teachers", handlers.TeacherHandler.CreateTeacher)
		admin.PUT("master/teachers/:uuid", handlers.TeacherHandler.CreateTeacher)
		admin.DELETE("master/teachers/:uuid", handlers.TeacherHandler.DeleteTeacher)

		admin.GET("master/students", handlers.StudentHandler.GetAllStudentsData)
		admin.GET("master/students/:uuid", handlers.StudentHandler.GetStudentData)
		admin.POST("master/students", handlers.StudentHandler.ImportStudentsData)
		admin.PUT("master/students/:uuid", handlers.StudentHandler.UpdateStudentData)
		admin.DELETE("master/students/:uuid", handlers.StudentHandler.DeleteStudent)

		admin.POST("master/classes", handlers.ClassHandler.CreateClass)
		admin.GET("master/classes/:uuid", handlers.StudentHandler.GetStudentData)
		admin.GET("master/classes", handlers.StudentHandler.ImportStudentsData)
		admin.PUT("master/classes/:uuid", handlers.StudentHandler.UpdateStudentData)
		admin.DELETE("master/classes/:uuid", handlers.StudentHandler.DeleteStudent)
	}

	// siswa := router.Group("/master/siswa").Use(middleware.IsValidJWT(db), middleware.IsAdmin())
	// {
	// 	siswa.GET("/", handlers.StudentHandler.GetAllStudentsData)
	// 	siswa.GET("/:uuid", handlers.StudentHandler.GetStudentData)
	// 	siswa.POST("/", middleware.IsExcelFile(), handlers.StudentHandler.ImportStudentsData)
	// 	siswa.PUT("/:uuid", middleware.ValidateStudentJSON(), handlers.StudentHandler.UpdateStudentData)
	// 	siswa.DELETE("/:uuid", handlers.StudentHandler.DeleteStudentData)
	// }

	return router
}
