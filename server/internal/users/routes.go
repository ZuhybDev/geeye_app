package users

import (
	"github.com/ZuhybDev/geeyeApp/middleware"
	"github.com/gofiber/fiber/v3"
)

func RegisterUserRoutes(api fiber.Router, h *Handler) {

	// user routes
	api.Get("/users", middleware.AuthMiddleware, h.GetListUsers)
	api.Post("/user", h.NewUser)
	api.Post("/user/login", h.Login)
	api.Patch("/user/:id", middleware.AuthMiddleware, h.UpdateUser)
	api.Delete("/user", middleware.AuthMiddleware, h.DeleteUser)
}
