package products

import (
	"github.com/ZuhybDev/geeyeApp/middleware"
	"github.com/gofiber/fiber/v3"
)

func RegisterProductRoutees(api fiber.Router, h *ProductsHandler) {

	api.Get("/feed/products", middleware.AuthMiddleware, h.UserFeedProducts)
	api.Get("/product", middleware.AuthMiddleware, h.GetAllUserProduct)
	api.Post("/product", middleware.AuthMiddleware, h.NewProducts)
	api.Patch("/product/:id", middleware.AuthMiddleware, h.UpdateProduct)
	api.Delete("/product/:id", middleware.AuthMiddleware, h.DeleteProduct)

}
