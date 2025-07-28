package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"simpleBE/internal"
)

func main() {
	// Initialize the database (connects and runs migrations/seeding)
	db := internal.InitDB()

	// Set up the service layer with the database
	svc := internal.NewService(db)

	// Create a new handler that uses the service
	handler := internal.NewHandler(svc)

	// Set up the Gin router
	r := gin.Default()

	// Define the GET endpoint for /api/hello
	r.GET("/api/hello", handler.GetHello)

	// Start the server on port 8080 and handle errors
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
