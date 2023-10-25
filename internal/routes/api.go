package routes

import (
	"github.com/gin-gonic/gin"
	customHTTP "github.com/iki-rumondor/init-golang-service/internal/adapter/http"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/middleware"
)

func StartServer(handlers *customHTTP.Handlers) *gin.Engine {
	router := gin.Default()

	router.MaxMultipartMemory = 10 << 20

	publicMaster := router.Group("/master")
	authorizationMaster := router.Group("/master")
	authorizationMaster.Use(middleware.ValidateHeader())
	{
		publicMaster.GET("/siswa", handlers.StudentHandler.GetAllStudentsData)
		publicMaster.GET("/siswa/:uuid", handlers.StudentHandler.GetStudentData)
		authorizationMaster.POST("/siswa", middleware.IsExcelFile(), handlers.StudentHandler.ImportStudentsData)
		authorizationMaster.PUT("/siswa/:uuid", middleware.ValidateStudentJSON(), handlers.StudentHandler.UpdateStudentData)
		authorizationMaster.DELETE("/siswa/:uuid", handlers.StudentHandler.DeleteStudentData)
	}

	auth := router.Group("/auth")
	{
		auth.POST("/login", middleware.ValidateLogin(), handlers.AuthHandler.Login)
		auth.POST("/verify-token", middleware.ValidateHeader(), handlers.AuthHandler.VerifyToken)
	}

	return router
}
