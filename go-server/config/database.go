package config

import (
	"fmt"
	"os"

	"github.com/Lavale1012/vector-db/go-server/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// Adjust the import path as necessary
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic("Error loading .env file")
	}
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	if connectionString == "" {
		panic("Database connection string not set in .env file")
	}
	fmt.Println("Connecting to the database with connection string")
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
	fmt.Println("Connected to the database successfully")
	DB = db
	DB.AutoMigrate(&models.VectorModel{})

}
