package routes

import (
	"github.com/gin-gonic/gin"
	customHTTP "github.com/iki-rumondor/init-golang-service/internal/adapter/http"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/middleware"
)

func StartServer(handler *customHTTP.StudentHandlers) *gin.Engine{
	router := gin.Default()

	router.MaxMultipartMemory = 10 << 20
	
	siswa := router.Group("/master")
	{
		siswa.GET("/endpoints")
		siswa.POST("/siswa", middleware.IsExcelFile(), handler.CreateStudent)
		siswa.GET("/endpoints:id")
		siswa.PUT("/endpoints/:id")
		siswa.DELETE("/endpoints/:id")
	}

	return router
}