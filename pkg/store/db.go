package store

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"music-library/pkg/model"
	"os"
)

var DB *gorm.DB

func InitDB() {
	// .env fayldan ma'lumotlarni o'qish
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"),
	)

	// connect data base
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	// Jadval yaratish
	err = DB.AutoMigrate(&model.Song{})
	if err != nil {
		log.Fatalf("Error migrating database: %s", err)
	}
}

func GetDB() *gorm.DB {
	return DB
}
