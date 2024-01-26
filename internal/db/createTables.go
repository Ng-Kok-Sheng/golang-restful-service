package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func CreateTables(pool *pgxpool.Pool) {
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
		log.Fatal("Something went wrong initializing tables: ", err)
	}

	log.Println("Tables created successfully")
}
