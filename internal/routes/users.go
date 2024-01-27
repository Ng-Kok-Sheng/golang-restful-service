package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang-restful/internal/db/users"
	"log"
	"net/http"
)

type UserHandlers struct {
	DbPool *pgxpool.Pool
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

}

func (app *UserHandlers) getUserById(c *gin.Context) {}

func (app *UserHandlers) createUser(c *gin.Context) {
	var user users.User
	if err := c.BindJSON(&user); err != nil {
		log.Printf("Something went wrong parsing user body: %s\n", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong parsing user body."})
		return
	}

	// TODO: To implement password hashing before inserting to DB
	if err := users.InsertUser(app.DbPool, &user); err != nil {
		log.Printf("Something went wrong inserting the user: %s\n", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong inserting user."})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "User successfully created", "id": user.ID})
}

func (app *UserHandlers) updateUser(c *gin.Context) {}

func (app *UserHandlers) deleteUser(c *gin.Context) {}
