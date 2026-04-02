package routes

import (
	"os"

	connection "github.com/ZuhybDev/geeyeApp/config"
	"github.com/ZuhybDev/geeyeApp/db"
	"github.com/ZuhybDev/geeyeApp/handlers"
	"github.com/ZuhybDev/geeyeApp/middleware"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {

	queries := db.New(connection.DBPool)

	secret := os.Getenv("JWT_SECRET")

	handler := &handlers.Handler{
		Query:     queries,
		JwtSecret: secret,
	}

	//Api group
	api := app.Group("/api")

	//user
	api.Get("/users", middleware.AuthMiddleware, handler.GetListUsers)
	api.Post("/user", handler.NewUser)
	api.Post("/user/login", handler.Login)
	api.Patch("/user/:id", middleware.AuthMiddleware, handler.UpdateUser)
	api.Delete("/user/", middleware.AuthMiddleware, handler.DeleteUser)

	//product
	api.Post("/restaurant", middleware.AuthMiddleware, handler.NewRestaurent)
	api.Delete("/restaurant/delete", middleware.AuthMiddleware, handler.DeleteRestaurant)
	api.Patch("/restaurant", middleware.AuthMiddleware, handler.UpdateRestaurant)
}
