package orders

import (
	"github.com/ZuhybDev/geeyeApp/middleware"
	"github.com/gofiber/fiber/v3"
)

func RegisterOrderRoutes(api fiber.Router, h *OrderHandler) {
	api.Post("/order", middleware.AuthMiddleware, h.CreatOrder)
}
