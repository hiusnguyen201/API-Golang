package models

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	dsn := os.Getenv("POSTGRES_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database!")
	}
}