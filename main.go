package main

import (
	"book-management-system/config"
	"book-management-system/models"
	"book-management-system/routes"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Initialize the database connection
	if err := config.InitDB(); err != nil {
		log.Fatal("Failed to initialize the database connection")
	}
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error on loading env file. err : %v", err)
		return
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Auto-migrate models
	if err := config.DB.AutoMigrate(&models.User{}, &models.Book{}, &models.BorrowRecord{}); err != nil {
		log.Fatal("Database migration failed")
	}

	// Set up the Gin router
	r := gin.Default()
	routes.RegisterRoutes(r)

	// Run the server in a separate goroutine
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	// Graceful shutdown setup
	go func() {
		log.Printf("Starting server on port :%s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server failed to start")
		}
	}()

	// Capture OS signals for graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive a shutdown signal
	<-stop
	log.Println("Shutting down server...")

	// Allow time for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown")
	}

	// If you want to close the underlying SQL connection pool managed by GORM
	if sqlDB, err := config.DB.DB(); err == nil {
		if err := sqlDB.Close(); err != nil {
			log.Fatal("Error closing database connection pool")
		}
	} else {
		log.Fatalf("Error getting underlying SQL DB")
	}
	log.Println("Application shutdown completed")
}
