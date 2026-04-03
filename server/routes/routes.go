package routes

import (
	connection "github.com/ZuhybDev/geeyeApp/config"
	"github.com/ZuhybDev/geeyeApp/db"
	env "github.com/ZuhybDev/geeyeApp/envConfig"
	"github.com/ZuhybDev/geeyeApp/internal"
	"github.com/ZuhybDev/geeyeApp/internal/restaurant"
	"github.com/ZuhybDev/geeyeApp/internal/users"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {

	queries := db.New(connection.DBPool)

	secret := env.ENV.JWTSecret

	appHandler := internal.App{
		Query:     queries,
		JwtSecret: secret,
	}

	// Api group
	api := app.Group("/api")

	// handler
	resHandler := restaurant.NewRestaurantHandler(&appHandler)
	userHandler := users.NewUserHandler(&appHandler)

	// pass the app and handler to route registers
	restaurant.RegisterRoutes(api, resHandler)
	users.RegisterUserRoutes(api, userHandler)

}
