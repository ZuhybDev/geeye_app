package routes

import (
	"github.com/ZuhybDev/geeyeApp/connection"
	env "github.com/ZuhybDev/geeyeApp/envConfig"
	"github.com/ZuhybDev/geeyeApp/internal/cars"
	delivers "github.com/ZuhybDev/geeyeApp/internal/deliver"
	"github.com/ZuhybDev/geeyeApp/internal/orders"
	products "github.com/ZuhybDev/geeyeApp/internal/product"
	"github.com/ZuhybDev/geeyeApp/internal/restaurant"
	"github.com/ZuhybDev/geeyeApp/internal/users"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {

	appHandler := *env.ENV

	dbPool := connection.DBPool
	// Api group
	api := app.Group("/api")

	// handler
	resHandler := restaurant.NewRestaurantHandler(&appHandler)
	userHandler := users.NewUserHandler(&appHandler)
	productHandler := products.NewProductHandler(&appHandler)
	carHandler := cars.NewCarHandler(&appHandler)
	deliverHandler := delivers.NewDeliverHandler(&appHandler)
	orderHandler := orders.NewOrderHandler(&appHandler, dbPool)

	// pass the app and handler to route registers
	restaurant.RegisterRoutes(api, resHandler)
	users.RegisterUserRoutes(api, userHandler)
	products.RegisterProductRoutees(api, productHandler)
	cars.RegsiterCarRoutes(api, carHandler)
	delivers.RegisterDeliverRoutes(api, deliverHandler)
	orders.RegisterOrderRoutes(api, orderHandler)
}
