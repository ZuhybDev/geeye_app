package products

import (
	"github.com/ZuhybDev/geeyeApp/middleware"
	"github.com/gofiber/fiber/v3"
)

func RegisterProductRoutees(api fiber.Router, h *ProductsHandler) {

	api.Post("/product", middleware.AuthMiddleware, h.NewProducts)

}
