package env

import (
	"os"

	"github.com/joho/godotenv"
)

/*
==> This file loads the Env in to memory and other file will reuse insaed of reloading
every you need to use Env env
*/
var ENV *Config

type Config struct {
	PORT      string
	DBUrl     string
	JWTSecret string
}

func Init() {

	godotenv.Load()

	ENV = &Config{
		PORT:      getEnv("PORT", "3000"),
		DBUrl:     getEnv("DATABASE_URL", ""),
		JWTSecret: getEnv("JWT_SECRET", "secret"),
	}
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}
