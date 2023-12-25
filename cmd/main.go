package main

import (
	"fmt"
	"log"
	"os"

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

	for _, model := range domain.RegisterModel() {
		if err := gormDB.Debug().AutoMigrate(model.Model); err != nil {
			log.Fatal(err.Error())
			return
		}
	}

	if err := gormDB.First(&domain.User{}).Error; err != nil {
		gormDB.Create(&domain.User{
			Nama:     "Admin",
			Username: "admin",
			Password: "123456",
		})
	}

	fmt.Println("Database migrated succeesfully")

	auth_repo := repository.NewAuthRepository(gormDB)
	auth_service := application.NewAuthService(auth_repo)
	auth_handler := customHTTP.NewAuthHandler(auth_service)

	student_repo := repository.NewStudentRepository(gormDB)
	student_service := application.NewStudentService(student_repo)
	student_handler := customHTTP.NewStudentHandler(student_service)

	teacher_repo := repository.NewTeacherRepository(gormDB)
	teacher_service := application.NewTeacherService(teacher_repo)
	teacher_handler := customHTTP.NewTeacherHandler(teacher_service)

	handlers := &customHTTP.Handlers{
		StudentHandler: student_handler,
		AuthHandler:    auth_handler,
		TeacherHandler: teacher_handler,
	}

	var PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	routes.StartServer(handlers).Run(":" + PORT)
}
