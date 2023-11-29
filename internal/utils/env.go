package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetDatabaseEnv() (map[string]interface{}, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err.Error())
	}

	dbHost := os.Getenv("PGHOST")
	dbPort := os.Getenv("PGPORT")
	dbUser := os.Getenv("PGUSER")
	dbPassword := os.Getenv("PGPASSWORD")
	dbName := os.Getenv("PGNAME")
	sslMode := os.Getenv("SSLMODE")

	var infoDB = map[string]interface{}{
		"host":     dbHost,
		"port":     dbPort,
		"user":     dbUser,
		"password": dbPassword,
		"name":     dbName,
		"sslmode":  sslMode,
	}

	return infoDB, nil
}
