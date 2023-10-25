package routes

import (
	"github.com/gin-gonic/gin"
	customHTTP "github.com/iki-rumondor/init-golang-service/internal/adapter/http"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/middleware"
)

func StartServer(handlers *customHTTP.Handlers) *gin.Engine {
	router := gin.Default()

	router.MaxMultipartMemory = 10 << 20

	siswa := router.Group("/master")
	{
		siswa.GET("/siswa", handlers.StudentHandler.GetAllStudentsData)
		siswa.POST("/siswa", middleware.IsExcelFile(), handlers.StudentHandler.CreateStudentsData)
		siswa.GET("/siswa/:uuid", handlers.StudentHandler.GetStudentData)
		siswa.PUT("/siswa/:uuid", middleware.ValidateStudentJSON(), handlers.StudentHandler.UpdateStudentData)
		siswa.DELETE("/siswa/:uuid", handlers.StudentHandler.DeleteStudentData)
	}

	auth := router.Group("/auth")
	{
		auth.POST("/register")
		auth.POST("/login")
	}

	return router
}
