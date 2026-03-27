package handlers

import (
	"log"
	"os"
	"time"

	"github.com/ZuhybDev/geeyeApp/db"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

// helper
type Handler struct {
	Query *db.Queries
}

// function get all users
func (h *Handler) GetListUsers(c fiber.Ctx) error {
	ctx := c.Context()

	users, err := h.Query.GetUserList(ctx)
	if err != nil {
		log.Println("Error fetching users:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
	}

	// 2. Fiber's c.JSON automatically sets the status to 200
	return c.JSON(users)
}

func (h *Handler) NewUser(c fiber.Ctx) error {

	ctx := c.Context()
	var user db.User

	secret := os.Getenv("JWT_SECRET")

	if err := c.Bind().Body(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	insertUser, err := h.Query.NewUser(ctx, db.NewUserParams{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		ImageUrl:    user.ImageUrl,
	})

	if err != nil {
		log.Fatal("failed to create new user: ", err)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tkn, err := token.SignedString([]byte(secret))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    insertUser,
		"token":   tkn,
	})

}
