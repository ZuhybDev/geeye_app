package handlers

import (
	"time"

	"github.com/ZuhybDev/geeyeApp/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

type LoginUser struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (h *Handler) Login(c fiber.Ctx) error {

	var lgnUser LoginUser

	if err := c.Bind().Body(&lgnUser); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if lgnUser.Email == "" || lgnUser.Password == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Email and password required"})
	}

	res, err := h.Query.UserLogin(c.Context(), lgnUser.Email)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	ok := utils.VerifyPassword(lgnUser.Password, res.Password)

	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// 4. Generate JWT using the REAL ID from the database
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    res.ID, // Use the ID returned by RETURNING *
		"name":  res.Name,
		"email": res.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	tkn, err := token.SignedString([]byte(h.JwtSecret))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	// Don't send the password back!
	res.Password = ""

	return c.Status(200).JSON(fiber.Map{
		"message": "Welcome Back!! " + res.Name,
		"user":    res,
		"token":   tkn,
	})
}
