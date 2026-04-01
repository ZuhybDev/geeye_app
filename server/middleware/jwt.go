package middleware

import (
	"github.com/ZuhybDev/geeyeApp/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c fiber.Ctx) error {

	// 1. Get cookie
	tkn := c.Cookies("token")
	if tkn == "" {
		return c.Status(401).SendString("Unauthorized")
	}

	// 2. Parse and validate
	token, err := jwt.ParseWithClaims(tkn, &utils.UserPayload{}, func(t *jwt.Token) (any, error) {
		return utils.JWTSecret, nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).SendString("Unauthorized")
	}

	// 3. Store user data in Fiber Locals for use in other handlers
	claims, ok := token.Claims.(*utils.UserPayload)
	if !ok {
		return c.Status(401).SendString("Invalid token claims")
	}
	c.Locals("user", claims)

	return c.Next()
}
