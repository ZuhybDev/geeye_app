package connection

import (
	"context"
	"log"

	env "github.com/ZuhybDev/geeyeApp/envConfig"
	"github.com/jackc/pgx/v5/pgxpool" // Use pgx v5
	"github.com/joho/godotenv"
)

// shared all across the app
var DBPool *pgxpool.Pool

func Connect() {
	//load the Env
	godotenv.Load()
	dbURL := env.ENV.DBUrl
	var err error
	// pgxpool handles the connection and the "Ping" automatically
	DBPool, err = pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	log.Println("Database connection established with pgxpool")
}
