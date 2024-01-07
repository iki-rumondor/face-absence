package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/database"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/registry"
	"github.com/iki-rumondor/init-golang-service/internal/routes"
)

func main() {
	gormDB, err := database.NewMysqlDB()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	for _, model := range registry.RegisterModels() {
		if err := gormDB.Debug().AutoMigrate(model.Model); err != nil {
			log.Fatal(err.Error())
			return
		}
	}

	handlers := registry.RegisterHandlers(gormDB)

	// for _, model := range registry.RegisterModels() {
	// 	if err := gormDB.Migrator().DropTable(model.Model); err != nil {
	// 		log.Fatal(err.Error())
	// 		return
	// 	}
	// }

	if err := gormDB.First(&domain.User{}).Error; err != nil {
		gormDB.Create(&domain.User{
			Nama:     "Admin",
			Username: "admin",
			Password: "123456",
		})
		gormDB.Create(&domain.Admin{
			Uuid:   uuid.NewString(),
			UserID: 1,
		})
	}

	// gormDB.Create(&domain.User{
	// 	Nama:     "Ilham",
	// 	Username: "ilham",
	// 	Password: "123456",
	// })
	// gormDB.Create(&domain.Student{
	// 	Uuid: uuid.NewString(),
	// 	NIS: "1232313123",
	// 	JK: "Laki-laki",
	// 	TempatLahir: "Gorontalo",
	// 	TanggalLahir: "2002-10-20",
	// 	Alamat: "Kota",
	// 	UserID: 3,
	// 	ClassID: 1,
	// })

	fmt.Println("Database migrated succeesfully")

	var PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	routes.StartServer(handlers).Run(":" + PORT)
}
