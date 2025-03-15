package main

import (
	"fmt"
	"log"
	"os"

	"mrktplc-auth/internal/authorized_domains"
	"mrktplc-auth/internal/handlers"
	"mrktplc-auth/internal/token"

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

	// get secret from environment variable or generate a new one
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "yes, this is the secret, please do something about it later"
	}

	// Set up routes
	authGroup := router.Group("/auth")
	{
		authHandler := handlers.NewAuthHandler(authorized_domains.NewAuthorizedDomains(), token.NewToken(secret))
		authGroup.POST("/enter", authHandler.Enter)
	}

	// Start server
	port := "8081"
	
	log.Printf("Server starting on port %s...", port)
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}