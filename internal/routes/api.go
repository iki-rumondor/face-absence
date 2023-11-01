package routes

import (
	"github.com/gin-gonic/gin"
	customHTTP "github.com/iki-rumondor/init-golang-service/internal/adapter/http"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/middleware"
)

func StartServer(handlers *customHTTP.Handlers) *gin.Engine {
	router := gin.Default()

	router.MaxMultipartMemory = 10 << 20

	public := router.Group("")
	{
		public.POST("/auth/login", handlers.AuthHandler.Login)
		public.GET("/auth/verify-token", middleware.ValidateHeader(), handlers.AuthHandler.VerifyToken)
	}

	siswa := router.Group("/master/siswa").Use(middleware.ValidateHeader(), middleware.IsAdmin())
	{
		siswa.GET("/", handlers.StudentHandler.GetAllStudentsData)
		siswa.GET("/:uuid", handlers.StudentHandler.GetStudentData)
		siswa.POST("/", middleware.IsExcelFile(), handlers.StudentHandler.ImportStudentsData)
		siswa.PUT("/:uuid", middleware.ValidateStudentJSON(), handlers.StudentHandler.UpdateStudentData)
		siswa.DELETE("/:uuid", handlers.StudentHandler.DeleteStudentData)
	}

	admin := router.Group("/master/admin").Use(middleware.ValidateHeader(), middleware.IsAdmin())
	{
		admin.POST("/", handlers.TeacherHandler.CreateTeacher)
		admin.GET("/", handlers.TeacherHandler.GetTeachers)
	}

	return router
}
