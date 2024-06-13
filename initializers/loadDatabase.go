package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectToDB() {
	LoadEnv()
	dsn := os.Getenv("DB_URI")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// You do not need to ping with gorm.DB. Instead, you can perform a query to check the connection.
	// For example, you can check if the connection works by querying any table or performing a raw SQL query.
	// Here is an example of checking if the connection is working by querying the current database version.
	var version string
	if err := DB.Raw("SELECT version()").Scan(&version).Error; err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}

	fmt.Printf("Connected to PostgreSQL database! Database version: %s\n", version)
}
