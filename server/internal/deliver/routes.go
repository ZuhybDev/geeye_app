package delivers

import (
	"github.com/ZuhybDev/geeyeApp/middleware"
	"github.com/gofiber/fiber/v3"
)

func RegisterDeliverRoutes(api fiber.Router, h *DeliverHandler) {
	api.Post("/deliver", h.NewDeliver)
	api.Patch("/deliver/:id", middleware.DeliverAuthMiddleware, h.UpdateDeliver)
	api.Delete("/deliver/:id", middleware.DeliverAuthMiddleware, h.DeleteDelivery)
	api.Get("/deliver/:id", middleware.DeliverAuthMiddleware, h.GetDeliver)
}
