package routes

import (
	"os"

	connection "github.com/ZuhybDev/geeyeApp/config"
	"github.com/ZuhybDev/geeyeApp/db"
	"github.com/ZuhybDev/geeyeApp/handlers"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {

	queries := db.New(connection.DBPool)

	secret := os.Getenv("JWT_SECRET")

	handler := &handlers.Handler{
		Query:     queries,
		JwtSecret: secret,
	}

	api := app.Group("/api")

	api.Get("/users", handler.GetListUsers)
	api.Post("/user", handler.NewUser)
	//user login
	api.Post("/login", handler.Login)
	api.Patch("/user/:id", handler.UpdateUser)
}
