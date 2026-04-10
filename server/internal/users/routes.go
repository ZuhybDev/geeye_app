package users

import (
	"github.com/ZuhybDev/geeyeApp/middleware"
	"github.com/gofiber/fiber/v3"
)

func RegisterUserRoutes(api fiber.Router, h *UserHandler) {

	// user routes
	api.Get("/users", middleware.AuthMiddleware, h.GetListUsers)
	api.Post("/user", h.NewUser)
	api.Post("/user/login", h.Login)
	api.Patch("/user/:id", middleware.AuthMiddleware, h.UpdateUser)
	api.Delete("/user", middleware.AuthMiddleware, h.DeleteUser)

	//user adresses
	api.Post("/user/address/new", middleware.AuthMiddleware, h.NewUserAddress)
	api.Patch("/user/address/:id", middleware.AuthMiddleware, h.UpdateAddress)
	api.Get("/user/addresses", middleware.AuthMiddleware, h.GetUserAddresses)
}
