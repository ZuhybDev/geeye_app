package env

import (
	"log"
	"os"

	"github.com/ZuhybDev/geeyeApp/connection"
	"github.com/ZuhybDev/geeyeApp/db"
	"github.com/joho/godotenv"
)

/*
==> This file loads the Env in to memory and other file will reuse insaed of reloading
every you need to use Env env
*/
var ENV *Config

type Config struct {
	Query            *db.Queries
	PORT             string
	JWTSecret        string
	DeliverJwtSecret string
	AdminJwtSecret   string
}

func Init() {

	// load the variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	queries := db.New(connection.DBPool)

	ENV = &Config{
		Query:            queries,
		PORT:             getEnv("PORT", "3000"),
		JWTSecret:        getEnv("JWT_SECRET", "secret"),
		DeliverJwtSecret: getEnv("DELIVER_JWT_SECRET", "deliverSecret"),
		AdminJwtSecret:   getEnv("ADMIN_JWT_SECRET", "admdinSecret"),
	}
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}
