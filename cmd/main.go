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
	gormDB, err := database.NewPostgresDB()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// for _, model := range domain.RegisterModel() {
	// 	if err := gormDB.Migrator().DropTable(model.Model); err != nil {
	// 		log.Fatal(err.Error())
	// 		return
	// 	}
	// }

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

	class_repo := repository.NewClassRepository(gormDB)
	class_service := application.NewClassService(class_repo)
	class_handler := customHTTP.NewClassHandler(class_service)

	subject_repo := repository.NewSubjectRepository(gormDB)
	subject_service := application.NewSubjectService(subject_repo)
	subject_handler := customHTTP.NewSubjectHandler(subject_service)

	sy_repo := repository.NewSchoolYearRepository(gormDB)
	sy_service := application.NewSchoolYearService(sy_repo)
	sy_handler := customHTTP.NewSchoolYearHandler(sy_service)

	schedule_repo := repository.NewScheduleRepository(gormDB)
	schedule_service := application.NewScheduleService(schedule_repo)
	schedule_handler := customHTTP.NewScheduleHandler(schedule_service)

	handlers := &customHTTP.Handlers{
		StudentHandler:    student_handler,
		AuthHandler:       auth_handler,
		TeacherHandler:    teacher_handler,
		ClassHandler:      class_handler,
		SubjectHandler:    subject_handler,
		SchoolYearHandler: sy_handler,
		ScheduleHandler:   schedule_handler,
	}

	var PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	routes.StartServer(handlers).Run(":" + PORT)
}
