package main

import (
	"log"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/database"
	customHTTP "github.com/iki-rumondor/init-golang-service/internal/adapter/http"
	"github.com/iki-rumondor/init-golang-service/internal/application"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"github.com/iki-rumondor/init-golang-service/internal/routes"
	"gorm.io/gorm"
)

func main() {
	gormDB, err := database.NewMysqlDB()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	
	// migration(gormDB)

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

func migration(db *gorm.DB) {
	
	db.Migrator().DropTable(&domain.Role{})
	db.Migrator().DropTable(&domain.Student{})
	db.Migrator().DropTable(&domain.User{})

	db.Migrator().CreateTable(&domain.Role{})
	db.Migrator().CreateTable(&domain.User{})
	db.Migrator().CreateTable(&domain.Student{})

	var roles = []domain.Role{
		{

			Name: "Siswa",
		},
		{
			Name: "Admin",
		},
	}
	db.Create(&roles)

	var user = domain.User{
		Uuid: "1",
		Nama: "Admin",
		Email: "admin@admin.com",
		Password: "123456",
		RoleID: 1,
	}
	db.Create(&user)
}
