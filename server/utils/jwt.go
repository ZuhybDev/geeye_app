package utils

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type UserPayload struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	RestaurentID string `json:"restaurent_id"`
	jwt.RegisteredClaims
}

func AuthMiddleware(c fiber.Ctx) error {

	fmt.Println("userPayload: ", UserPayload{})
	// 1. Get cookie
	cookie := c.Cookies("token")
	if cookie == "" {
		return c.Status(401).SendString("Unauthorized")
	}

	// 2. Parse and validate
	token, err := jwt.ParseWithClaims(cookie, &UserPayload{}, func(t *jwt.Token) (any, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).SendString("Invalid session")
	}

	// 3. Store user data in Fiber Locals for use in other handlers
	claims := token.Claims.(*UserPayload)
	c.Locals("user", claims)

	return c.Next()
}
