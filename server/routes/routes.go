package routes

import (
	connection "github.com/ZuhybDev/geeyeApp/config"
	"github.com/ZuhybDev/geeyeApp/db"
	"github.com/ZuhybDev/geeyeApp/handlers"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {

	queries := db.New(connection.DBPool)

	handler := &handlers.Handler{
		Query: queries,
	}

	api := app.Group("/api")

	api.Get("/users", handler.GetListUsers)
	api.Post("/user", handler.NewUser)

}
