package cars

import (
	"github.com/ZuhybDev/geeyeApp/middleware"
	"github.com/gofiber/fiber/v3"
)

func RegsiterCarRoutes(api fiber.Router, h *carHandler) {

	api.Get("/car/:id", middleware.AuthMiddleware, h.GetCarById)
	api.Post("/car", middleware.AuthMiddleware, h.CreateNewCar)
	api.Patch("/car/:id", middleware.AuthMiddleware, h.UpdateCar)
	api.Delete("/car/:id", middleware.AuthMiddleware, h.DeleteCar)
}
