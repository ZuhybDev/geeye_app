package restaurant

import (
	"github.com/ZuhybDev/geeyeApp/middleware"
	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(api fiber.Router, h *ResHandler) {

	// //restaurant
	api.Post("/restaurant", middleware.AuthMiddleware, h.NewRestaurent)
	api.Patch("/restaurant/update", middleware.AuthMiddleware, h.UpdateRestaurant)
	api.Delete("/restaurant/delete", middleware.AuthMiddleware, h.DeleteRestaurant)

	// Res Addresses
	api.Post("/restaurant/address", middleware.AuthMiddleware, h.CreateResAddress)
	// get addresses
	api.Get("/restaurant/addresses", middleware.AuthMiddleware, h.GetAdderessById)
}
