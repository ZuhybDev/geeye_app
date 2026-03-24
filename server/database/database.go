package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	// dbURL := os.Getenv("DATABASE_URL")
	dbURL := "postgres://postgres@127.0.0.1:5432/delivery_db?sslmode=disable"
	var err error

	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	// Check if connection is actually alive
	err = DB.Ping()
	if err != nil {
		log.Fatal("Database is unreachable:", err)
	}

	log.Println("Database connection established")
}
