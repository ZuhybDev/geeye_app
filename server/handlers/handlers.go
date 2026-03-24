package handlers

import "github.com/gofiber/fiber/v3"

func HelloRoute(c fiber.Ctx) error {
	return c.SendString("Hello from handlers")
}
