package middleware

import (
	"fmt"
	"os"

	env "github.com/ZuhybDev/geeyeApp/envConfig"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

type UserPayload struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

var secret = os.Getenv("JWTSecret")

func AuthMiddleware(c fiber.Ctx) error {

	// 1. Get cookie
	tkn := c.Cookies("token")
	if tkn == "" {
		return c.Status(404).SendString("Unauthorized token is missing")
	}

	fmt.Printf("Token from ENV file: %s\n", env.ENV.JWTSecret)
	// 2. Parse and validate
	token, err := jwt.ParseWithClaims(tkn, &UserPayload{}, func(t *jwt.Token) (any, error) {
		return env.ENV.JWTSecret, nil
	})

	if err != nil || !token.Valid {
		return c.Status(404).SendString("Unauthorized invalid token")
	}

	// 3. Store user data in Fiber Locals for use in other handlers
	claims, ok := token.Claims.(*UserPayload)
	if !ok {
		return c.Status(401).SendString("Invalid token claims")
	}

	c.Locals("user", claims)

	return c.Next()
}
