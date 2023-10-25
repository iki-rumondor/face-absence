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
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	var roles = []domain.Role{
		{

			Name: "Siswa",
		},
		{
			Name: "Admin",
		},
	}

	gormDB.Debug().AutoMigrate(&domain.Role{})
	gormDB.Create(&roles)
	gormDB.Debug().AutoMigrate(&domain.Student{})
	gormDB.Debug().AutoMigrate(&domain.User{})

	student_repo := repository.NewStudentRepository(gormDB)
	student_service := application.NewStudentService(student_repo)
	student_handler := customHTTP.NewStudentHandler(student_service)

	auth_repo := repository.NewAuthRepository(gormDB)
	auth_service := application.NewAuthService(auth_repo)
	auth_handler := customHTTP.NewAuthHandler(auth_service)

	handlers := &customHTTP.Handlers{
		StudentHandler: student_handler,
		AuthHandler:    auth_handler,
	}

	var PORT = ":8082"
	routes.StartServer(handlers).Run(PORT)
}
