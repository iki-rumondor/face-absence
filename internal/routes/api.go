package routes

import (
	"github.com/gin-gonic/gin"
	customHTTP "github.com/iki-rumondor/init-golang-service/internal/adapter/http"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/middleware"
)

func StartServer(handler *customHTTP.StudentHandlers) *gin.Engine {
	router := gin.Default()

	router.MaxMultipartMemory = 10 << 20

	siswa := router.Group("/master")
	{
		siswa.GET("/siswa", handler.GetAllStudentsData)
		siswa.POST("/siswa", middleware.IsExcelFile(), handler.CreateStudentsData)
		siswa.GET("/siswa/:uuid", handler.GetStudentData)
		siswa.PUT("/siswa/:uuid", middleware.ValidateStudentJSON(), handler.UpdateStudentData)
		siswa.DELETE("/siswa/:uuid", handler.DeleteStudentData)
	}

	return router
}
