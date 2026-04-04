package restaurant

import (
	"github.com/ZuhybDev/geeyeApp/middleware"
	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(api fiber.Router, h *ResHandler) {

	// //restaurant
	api.Post("/restaurant", middleware.AuthMiddleware, h.NewRestaurent)
	api.Delete("/restaurant/delete", middleware.AuthMiddleware, h.DeleteRestaurant)
	api.Patch("/restaurant", middleware.AuthMiddleware, h.UpdateRestaurant)
}
