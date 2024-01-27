package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetPostgresConnectionPool() *pgxpool.Pool {
	// To use env vars
	postgresUri := "postgres://admin:root@localhost:5432/test_db"

	dbPool, err := pgxpool.New(context.Background(), postgresUri)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	return dbPool
}
