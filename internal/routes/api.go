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

	master := router.Group("/master").Use(middleware.ValidateHeader(), middleware.IsAdmin())
	{
		master.GET("/siswa", handlers.StudentHandler.GetAllStudentsData)
		master.GET("/siswa/:uuid", handlers.StudentHandler.GetStudentData)
		master.POST("/siswa", middleware.IsExcelFile(), handlers.StudentHandler.ImportStudentsData)
		master.PUT("/siswa/:uuid", middleware.ValidateStudentJSON(), handlers.StudentHandler.UpdateStudentData)
		master.DELETE("/siswa/:uuid", handlers.StudentHandler.DeleteStudentData)
	}

	return router
}
