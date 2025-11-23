package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/darrkeer/avito-tech-test-task/handlers"
	"github.com/darrkeer/avito-tech-test-task/repository"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	log.Print("START")

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL not setup")
	}
	log.Print("DATABASE FOUND")

	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()
	log.Print("DATABASE CONNECTION ESTABLISHED")

	repo := repository.New(db)
	handler := handlers.New(repo)
	handler.Start()

	log.Println("STARTING SERVICE IN PORT: 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("HTTP SERVER STOPPED: %v", err)
	}
}
