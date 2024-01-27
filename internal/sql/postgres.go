// Package sql provides a streamlined interface for initializing connection pools with PostgreSQL and includes helper
// functions for building and executing SQL queries. This package aims to simplify the interaction with PostgreSQL
// databases by abstracting the intricacies of database connection management and query construction.
package sql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

// GetPostgresConnectionPool is used to get a connection pool to postgres server.
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
