package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		log.Fatal("Environment variable DATABASE_DSN not set")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true, Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")
}

func Close() {
	db, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
	log.Println("Disconnected from database")
}
