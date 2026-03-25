package connection

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool" // Use pgx v5
	"github.com/joho/godotenv"
)

// Use Pool instead of *sql.DB for better performance with Fiber
var DBPool *pgxpool.Pool

func Connect() {
	godotenv.Load()
	dbURL := os.Getenv("DATABASE_URL")

	var err error
	// pgxpool handles the connection and the "Ping" automatically
	DBPool, err = pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	log.Println("Database connection established with pgxpool")
}
