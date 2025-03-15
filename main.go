package main

import (
	"fmt"
	"log"

	"mrktplc-auth/internal/authorized_domains"
	"mrktplc-auth/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Create Gin router
	router := gin.Default()

	// Set up routes
	authGroup := router.Group("/auth")
	{
		authHandler := handlers.NewAuthHandler(authorized_domains.NewAuthorizedDomains())
		authGroup.POST("/enter", authHandler.Enter)
	}

	// Start server
	port := "8081"
	
	log.Printf("Server starting on port %s...", port)
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}