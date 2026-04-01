package utils

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret []byte

type UserPayload struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	RestaurantID string `json:"restaurent_id"`
	jwt.RegisteredClaims
}

func LoadConfig() {
	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	JWTSecret = []byte(secret)
}
