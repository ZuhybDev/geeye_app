package middleware

import (
	"os" // New in v3

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v3"
)

// Initialize the middleware once
var NewJWTAuth = jwtware.New(jwtware.Config{
	SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))}, // Check if you meant "JWT_SECRET" instead of "WT_SECRET"
	ContextKey: "jwt",
	ErrorHandler: func(c fiber.Ctx, err error) error {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	},
})
