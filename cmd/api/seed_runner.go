package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"auth-service-2.0/internal/db"
	_ "github.com/lib/pq"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide the number of entries to seed (e.g., go run seed_runner.go 10)")
	}
	count, err := strconv.Atoi(os.Args[1])
	if err != nil || count <= 0 {
		log.Fatal("Invalid count provided. Must be a positive integer.")
	}

	connStr := os.Getenv("DB_ADDR")
	fmt.Println("Connecting to database:", connStr)
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer func(dbConn *sql.DB) {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(dbConn)

	if err := db.Seed(dbConn, count); err != nil {
		log.Fatal("Seeding failed:", err)
	}
	log.Println("Seeding successful.")
}
