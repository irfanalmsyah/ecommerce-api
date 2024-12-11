package database

import (
	"log"
	"os"
	"strings"

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

func IsTableEmpty(model interface{}) bool {
	var count int64
	if err := DB.Model(model).Count(&count).Error; err != nil {
		log.Printf("Error counting records: %v\n", err)
		return false
	}
	return count == 0
}

func SeedDatabase(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	sqlCommands := string(data)
	statements := strings.Split(sqlCommands, ";")

	for _, stmt := range statements {
		trimmed := strings.TrimSpace(stmt)
		if trimmed == "" {
			continue
		}
		if err := DB.Exec(trimmed).Error; err != nil {
			return err
		}
	}

	log.Println("Database seeded successfully.")
	return nil
}
