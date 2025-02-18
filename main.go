package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"hermes/config"
	"hermes/handlers"
	"hermes/redis"
)

// setupRouter initializes all API routes
func setupRouter() *gin.Engine {
	r := gin.Default()

	// Health Check
	r.GET("/health", handlers.HealthCheck)

	// Authentication Routes
	r.POST("/register", handlers.Register)
	r.POST("/renew", handlers.Renew)

	// Messaging Routes
	r.POST("/publish", handlers.Publish)
	r.POST("/subscribe", handlers.Subscribe)

	// Redis Streams Routes
	r.POST("/streams", handlers.CreateStream)
	r.GET("/streams/:stream", handlers.ReadStream)

	// Metrics & Logs
	r.GET("/metrics", handlers.Metrics)
	r.GET("/logs", handlers.GetLogs)

	return r
}

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using default values")
	}

	// Initialize Redis
	redis.InitRedis()

	// Setup Gin mode (default: debug)
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = "release" // Default to release for production
	}
	gin.SetMode(ginMode)

	// Initialize router
	r := setupRouter()

	// Start server
	port := config.GetEnv("PORT", "8080")
	fmt.Printf("🚀 Hermes API Server Running on :%s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
