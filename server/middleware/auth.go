package middleware

import (
	"fmt"

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

func AuthMiddleware(c fiber.Ctx) error {

	// 1. Get cookie
	tkn := c.Cookies("token")
	if tkn == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"message": "Token is missing",
		})
	}

	// 2. Parse and validate
	token, err := jwt.ParseWithClaims(tkn, &UserPayload{}, func(t *jwt.Token) (any, error) {
		// Verify the signing method to prevent algorithm confusion attacks
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(env.ENV.JWTSecret), nil
	},
	)

	if err != nil || !token.Valid {
		// fmt.Printf("jwt Error: %v\n", err)
		return c.Status(401).SendString("Unauthorized invalid token")
	}

	// 3. Store user data in Fiber Locals for use in other handlers
	claims, ok := token.Claims.(*UserPayload)
	if !ok {
		return c.Status(401).SendString("Invalid token claims")
	}

	c.Locals("user", claims)

	return c.Next()
}
func DeliverAuthMiddleware(c fiber.Ctx) error {

	// 1. Get cookie
	tkn := c.Cookies("token")
	if tkn == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"message": "Token is missing",
		})
	}

	// 2. Parse and validate
	token, err := jwt.ParseWithClaims(tkn, &UserPayload{}, func(t *jwt.Token) (any, error) {
		// Verify the signing method to prevent algorithm confusion attacks
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(env.ENV.AdminJwtSecret), nil
	},
	)

	if err != nil || !token.Valid {
		// fmt.Printf("jwt Error: %v\n", err)
		return c.Status(401).SendString("Unauthorized invalid token")
	}

	// 3. Store user data in Fiber Locals for use in other handlers
	claims, ok := token.Claims.(*UserPayload)
	if !ok {
		return c.Status(401).SendString("Invalid token claims")
	}

	c.Locals("deliver", claims)

	return c.Next()
}
