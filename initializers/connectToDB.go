package initializers

import (
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



var DB *gorm.DB

func ConnectToDB() {
	dsn := os.Getenv("DB")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to conned to DB")
	}

	DB = db
}