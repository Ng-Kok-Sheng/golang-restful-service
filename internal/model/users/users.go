package users

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"strconv"
	"strings"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserUpdateRequest struct {
	Password *string `json:"password"`
	Email    *string `json:"email"`
}

func CreateTable(pool *pgxpool.Pool) {
	log.Println("Creating tables...")
	ctx := context.Background()

	installUuidOssp := `
	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
	`

	_, err := pool.Exec(ctx, installUuidOssp)
	if err != nil {
		log.Fatal("Something went wrong installing the UUID module: ", err)
	}

	createSQLTable := `
	CREATE TABLE IF NOT EXISTS users (
	    id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
	    username VARCHAR(20) UNIQUE NOT NULL,
	    password VARCHAR(100) UNIQUE NOT NULL,
	    email VARCHAR(100) UNIQUE NOT NULL,
	    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err = pool.Exec(ctx, createSQLTable)
	if err != nil {
		log.Fatal("Something went wrong initializing table user: ", err)
	}

	log.Println("Table user created successfully")
}

func InsertUser(dbPool *pgxpool.Pool, user *User) error {
	query := `
	INSERT into users (username, password, email) VALUES ($1, $2, $3) RETURNING id
	`

	return dbPool.QueryRow(context.Background(), query, user.Username, user.Password, user.Email).Scan(&user.ID)
}

func GetAllUsers(dbPool *pgxpool.Pool) (users []User, err error) {
	query := `SELECT * FROM users`

	rows, err := dbPool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func GetUserByUsernameAndPassword(dbPool *pgxpool.Pool, username string, password string) (user User, err error) {
	query := `SELECT * FROM users WHERE username = $1 AND password = $2`

	err = dbPool.QueryRow(context.Background(), query, username, password).Scan(
		&user.ID, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt,
	)

	return
}

func UpdateUser(dbPool *pgxpool.Pool, id string, updateRequest UserUpdateRequest) (err error) {
	query := `UPDATE users SET `
	var args []interface{}
	argId := 1

	if updateRequest.Password != nil {
		query += "password = $" + strconv.Itoa(argId) + ", "
		args = append(args, updateRequest.Password)
		argId++
	}

	if updateRequest.Email != nil {
		query += "email = $" + strconv.Itoa(argId) + ", "
		args = append(args, updateRequest.Email)
		argId++
	}

	query = strings.TrimSuffix(query, ", ")
	query += " WHERE id = $" + strconv.Itoa(argId)
	args = append(args, id)

	_, err = dbPool.Exec(context.Background(), query, args...)
	return
}

func DeleteUser(dbPool *pgxpool.Pool) (user User, err error) {
	return
}
