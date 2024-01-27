package main

import (
	"fmt"
	"log"
	"os"

	"github.com/iki-rumondor/init-golang-service/cmd/migrate"
	"github.com/iki-rumondor/init-golang-service/internal/adapter/database"
	"github.com/iki-rumondor/init-golang-service/internal/registry"
	"github.com/iki-rumondor/init-golang-service/internal/routes"
	"github.com/iki-rumondor/init-golang-service/internal/utils"
)

func main() {
	gormDB, err := database.NewMysqlDB()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	if len(os.Args)-1 > 0 {
		migrate.ReadTerminal(gormDB, os.Args)
		return
	}

	repo := registry.RegisterRepositories(gormDB)
	services := registry.RegisterServices(repo)
	handlers := registry.RegisterHandlers(services)

	fmt.Println("Database migrated succeesfully")

	var PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	utils.InitCustomValidation()

	routes.StartServer(handlers).Run(":" + PORT)
}
