package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConnection := os.Getenv("POSTGRESQL_CONNECTION")

	db, err := gorm.Open(postgres.Open(dbConnection), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
