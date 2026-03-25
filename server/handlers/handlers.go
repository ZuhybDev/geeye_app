package handlers

import (
	"log"

	connection "github.com/ZuhybDev/geeyeApp/config"
	"github.com/ZuhybDev/geeyeApp/db"
	"github.com/gofiber/fiber/v3"
)

func HelloRoute(c fiber.Ctx) error {
	return c.SendString("Hello from handlers")
}

// function get all users

func GetListUsers(c fiber.Ctx) error {
	ctx := c.Context()

	queries := db.New(connection.DBPool)

	// 3. List all users
	users, err := queries.GetUserList(ctx)
	if err != nil {
		log.Println("Error fetching users:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
	}

	// 4. Fiber's c.JSON automatically sets the status to 200
	return c.JSON(users)
}
