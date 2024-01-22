package database

import (
	"fmt"

	"github.com/iki-rumondor/init-golang-service/internal/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlDB() (*gorm.DB, error) {
	env, err := utils.GetDatabaseEnv()
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", env["user"], env["password"], env["host"], env["port"], env["name"])
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return gormDB, nil

}
