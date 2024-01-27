package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"golang-restful/internal/db"
	"golang-restful/internal/routes"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	router := gin.Default()
	postgresPool := db.GetPostgresConnectionPool()
	defer postgresPool.Close()

	// Register user routes
	routes.RegisterUserRoutes(router, postgresPool)

	// Run with goroutine so it doesn't block
	go func() {
		if err := router.Run("localhost:8080"); err != nil {
			log.Fatalf("Failed to run server: %v\n", err)
		}
	}()

	// Channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	// Catch CTRL+C and Kubernetes pod shutdown etc. (SIGINT / SIGTERM)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until signal is received
	<-quit
	log.Println("Shutting down server...")

	// Wait 5 seconds
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Server gracefully stopped")
}
