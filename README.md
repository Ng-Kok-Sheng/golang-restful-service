# GoLang Restful Server (To come up with a better name)
***
Personal side project to explore writing a restful server with other programming languages and tools. Project is a
simple user API with authentication and common CRUD operations.

- GoLang
- Testing in GoLang
- PostgreSQL
- DB queries without ORM
- Redis as cache

## Packages Used
***
- gin
- pgx

## Prerequisites
***
- Using GoLang 1.19
- Ensure docker daemon is running
- Ensure docker compose is available

## Environment Variables
***
Ensure that these environment variables are exported.
- SERVER_PORT
```bash
export SERVER_PORT=8080
```
- POSTGRES_URI
```bash
export POSTGRES_URI=postgres://<username>:<password>@localhost:5432/test_db
```

## Instructions to start
***
1. Open terminal and run 
```bash
docker compose up -d
```

2. Open another terminal and run
```bash
go run main.go
```

pgAdmin4 can be access through [localhost:5050](localhost:5050).