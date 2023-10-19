package main

import (
	"log"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/database"
	customHTTP "github.com/iki-rumondor/init-golang-service/internal/adapter/http"
	"github.com/iki-rumondor/init-golang-service/internal/application"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"github.com/iki-rumondor/init-golang-service/internal/routes"
)

func main() {
	gormDB, err := database.NewMysqlDB()
	if err != nil{
		log.Fatal(err.Error())
		return
	}
	
	gormDB.Debug().AutoMigrate(&domain.Student{})

	repo := repository.NewStudentRepository(gormDB)
	service := application.NewStudentService(repo)
	handler := customHTTP.NewStudentHandler(service)

	var PORT = ":8082"
	routes.StartServer(handler).Run(PORT)
}