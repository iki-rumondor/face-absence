package migrate

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/iki-rumondor/init-golang-service/internal/domain"
	"github.com/iki-rumondor/init-golang-service/internal/registry"
	"gorm.io/gorm"
)

func ReadTerminal(db *gorm.DB, args []string) {
	switch {
	case args[1] == "fresh":
		if err := freshDatabase(db); err != nil {
			log.Fatal(err.Error())
		}
	default:
		fmt.Println("Hello")
	}
}

func freshDatabase(db *gorm.DB) error {
	for _, model := range registry.RegisterModels() {
		if err := db.Migrator().DropTable(model.Model); err != nil {
			return err
		}
	}
	for _, model := range registry.RegisterModels() {
		if err := db.Debug().AutoMigrate(model.Model); err != nil {
			return err
		}
	}

	db.Create(&domain.User{
		Nama:     "Admin",
		Username: "admin",
		Password: "123456",
	})
	db.Create(&domain.Admin{
		Uuid:   uuid.NewString(),
		UserID: 1,
	})

	return nil
}
