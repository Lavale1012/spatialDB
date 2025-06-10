package main

import (
	"log"

	"github.com/Lavale1012/vector-db/go-server/api"
	"github.com/Lavale1012/vector-db/go-server/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	PORT := "8080"
	router := gin.Default()

	router.Use(gin.Recovery()) // Use recovery middleware to handle panics
	config.ConnectDB()
	router.Use(gin.Logger()) // Use logger middleware to log requests
	api.ApiRoutes(router)

	if err := router.Run("localhost:" + PORT); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	log.Println("Server running on http://localhost:" + PORT)
}
