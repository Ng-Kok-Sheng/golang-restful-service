package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang-restful/internal/db"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	dbPool *pgxpool.Pool
}

type user struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func main() {
	app := App{
		dbPool: db.GetPostgresConnectionPool(),
	}

	// Schedule the pool to close when main() returns
	defer app.dbPool.Close()

	router := gin.Default()

	router.GET("/users", app.getUsersHandler)
	//router.GET("/user/:id", getUserById)
	//router.POST("/user", createUser)

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

func (app *App) getUsersHandler(c *gin.Context) {
	var results string
	err := app.dbPool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&results)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No users found."})
		return
	}

	c.IndentedJSON(http.StatusOK, results)
}
