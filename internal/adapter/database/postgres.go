package database

import (
	"fmt"

	"github.com/iki-rumondor/init-golang-service/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB() (*gorm.DB, error) {

	env, err := utils.GetPostgresDeployEnv()
	if err != nil{
		return nil, err
	}
	strConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", env["host"], env["port"], env["user"], env["password"], env["name"])

	fmt.Println(strConn)
	gormDB, err := gorm.Open(postgres.Open(strConn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return gormDB, nil
}
