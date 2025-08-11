// Manages application configuration

package internal

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// Define the PostgreSQL DSN using service name 'db_dev' for Docker Compose networking
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		if os.Getenv("APP_ENV") == "dev" {
			dsn = "host=localhost port=5434 user=postgres password=postgres dbname=devdb sslmode=disable"
		} else {
			log.Fatal("DATABASE_URL not set")
		}
	}

	var db *gorm.DB
	var err error

	// Retry logic: attempt to connect up to 10 times, sleeping between each try
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Println("Waiting for database to be ready...")
		time.Sleep(2 * time.Second)
	}

	// Fatal error if unable to connect after retries
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	// Automatically create the messages table if it doesn't exist
	if err := db.AutoMigrate(&Message{}); err != nil {
		log.Fatal("Failed to migrate DB:", err)
	}

	// Check if a seed message already exists; insert "hello world" if not
	var count int64
	db.Model(&Message{}).Where("id = ?", 1).Count(&count)
	if count == 0 {
		db.Create(&Message{ID: 1, Text: "hello world"})
	}

	return db
}
