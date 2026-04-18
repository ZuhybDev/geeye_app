package delivers

import (
	"github.com/gofiber/fiber/v3"
)

func RegisterDeliverRoutes(api fiber.Router, h *DeliverHandler) {
	api.Post("/deliver", h.NewDeliver)
}
