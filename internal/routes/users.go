package routes

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
	"time"
)

type UserHandlers struct {
	DbPool *pgxpool.Pool
}

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func RegisterUserRoutes(router *gin.Engine, dbPool *pgxpool.Pool) {
	userHandler := UserHandlers{
		DbPool: dbPool,
	}

	router.GET("/users", userHandler.getAllUsers)
	router.GET("/user/:id", userHandler.getUserById)
	router.POST("/user", userHandler.createUser)
	router.PUT("/user/:id", userHandler.updateUser)
	router.DELETE("/user/:id", userHandler.deleteUser)
}

func (app *UserHandlers) getAllUsers(c *gin.Context) {
	var results string
	err := app.DbPool.QueryRow(context.Background(), "select * from users'").Scan(&results)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No users found."})
		return
	}

	c.IndentedJSON(http.StatusOK, results)
}

func (app *UserHandlers) getUserById(c *gin.Context) {}

func (app *UserHandlers) createUser(c *gin.Context) {}

func (app *UserHandlers) updateUser(c *gin.Context) {}

func (app *UserHandlers) deleteUser(c *gin.Context) {}
