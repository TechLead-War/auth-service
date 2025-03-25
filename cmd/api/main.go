package main

import (
	"database/sql"
	"log"

	"auth-service-2.0/internal/db"
	"auth-service-2.0/internal/repository"
	"auth-service-2.0/internal/resources"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg := config{
		AppAddr: resources.GetEnv("ADDR", ":8800"),
		DBConfig: dbConfig{
			DBAddr:       resources.GetEnv("DB_ADDR", "postgres://postgres_user:postgres_pass@localhost:5432/postgres_db?sslmode=disable"),
			MaxOpenConns: resources.GetEnvAsInt("DB_MAX_OPEN_CONNS", 30),
			MaxIdleConns: resources.GetEnvAsInt("DB_MAX_IDLE_CONNS", 30),
			MaxIdleTime:  resources.GetEnv("DB_MAX_IDLE_TIME", "15m"),
		},
		Env: resources.GetEnv("ENV", "development"),
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
