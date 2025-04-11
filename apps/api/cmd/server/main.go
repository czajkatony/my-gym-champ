package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/czajkatony/my-gym-champ/api/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	if err := database.Initialize(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	// Initialize router
	router := gin.Default()

	// Enable CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Health check route
	router.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"service": "my-gym-champ-api",
		})
	})

	// Setup API routes
	v1 := router.Group("/api/v1")
	{
		// Member routes
		members := v1.Group("/members")
		{
			members.GET("", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "List all members"})
			})
			members.GET("/:id", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "Get member details"})
			})
		}

		// Gym routes
		gyms := v1.Group("/gyms")
		{
			gyms.GET("", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "List all gyms"})
			})
			gyms.GET("/:id", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "Get gym details"})
			})
		}

		// Workout routes
		workouts := v1.Group("/workouts")
		{
			workouts.GET("", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "List all workouts"})
			})
			workouts.POST("", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "Create a new workout"})
			})
		}

		// Event routes
		events := v1.Group("/events")
		{
			events.GET("", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "List all events"})
			})
			events.POST("", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "Create a new event"})
			})
		}
	}

	// Set up server with graceful shutdown
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Start the server in a goroutine
	go func() {
		log.Println("Server starting on port 8080...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Give outstanding requests a timeout of 5 seconds to complete
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited properly")
}
