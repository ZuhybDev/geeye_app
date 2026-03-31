package handlers

import (
	"fmt"
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

	claims := utils.UserPayload{
		ID:           res.ID.String(),
		Name:         res.Name,
		Email:        res.Email,
		RestaurentID: res.RestaurantID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
			Issuer:    "geeye-app",
		},
	}

	// 4. Generate JWT using the REAL ID from the database
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tkn, err := token.SignedString(h.JwtSecret)

	fmt.Println(token)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	// 4. Set the Cookie
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tkn,
		Expires:  time.Now().Add(72 * time.Hour),
		HTTPOnly: true,  // Important: Prevents JS from stealing the token
		Secure:   false, //TODO Set to true in production with HTTPS
		SameSite: "Lax",
	})

	// Don't send the password back!
	res.Password = ""

	return c.Status(200).JSON(fiber.Map{
		"message": "Welcome Back!! " + res.Name,
		"user":    res,
	})
}
