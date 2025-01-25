package main

import (
	"database/sql"
	"log"

	"auth-service-2.0/internal/db"
	"auth-service-2.0/internal/envs"
	"auth-service-2.0/internal/repository"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("internal/envs/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg := config{
		AppAddr: envs.GetEnv("ADDR", ":8800"),
		DBConfig: dbConfig{
			DBAddr:       envs.GetEnv("DB_ADDR", "postgres://postgres_user:postgres_pass@localhost:5432/postgres_db?sslmode=disable"),
			MaxOpenConns: envs.GetEnvAsInt("DB_MAX_OPEN_CONNS", 30),
			MaxIdleConns: envs.GetEnvAsInt("DB_MAX_IDLE_CONNS", 30),
			MaxIdleTime:  envs.GetEnv("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	database, err := db.NewDBConnection(
		cfg.DBConfig.DBAddr,
		cfg.DBConfig.MaxOpenConns,
		cfg.DBConfig.MaxIdleConns,
		cfg.DBConfig.MaxIdleTime)

	if err != nil {
		log.Fatalf("Error creating database connection: %v", err)
	}

	// This allows handling any errors that occur when closing the database connection.
	// If an error occurs, it logs the error and terminates the program.
	// This is a closure is a function value that references variables from outside its body.
	// Closures are often used to create functions that have some state or to pass functions as arguments.
	defer func(database *sql.DB) {
		err := database.Close()
		if err != nil {
			log.Fatalf("Error closing database connection: %v", err)
		}
	}(database)

	log.Println("Database connection established")

	store := repository.NewStorage(database)

	app := app{
		AppAddr: cfg,
		Store:   store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
