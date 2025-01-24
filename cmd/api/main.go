package main

import (
	"log"

	"auth-service-2.0/internal/envs"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("internal/envs/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg := config{
		addr: envs.GetEnv("ADDR", ":8800"),
	}
	app := app{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
