package handlers

import "github.com/gofiber/fiber/v3"

func (h *Handler) NewRestaurent(c fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "Hello restaurant",
	})
}
