package main

import (
	"log"

	"github.com/ZuhybDev/geeyeApp/connection"
	env "github.com/ZuhybDev/geeyeApp/envConfig"
	"github.com/ZuhybDev/geeyeApp/routes"
	"github.com/gofiber/fiber/v3"
)

func main() {

	// Connect once
	connection.Connect()

	if connection.DBPool == nil {
		log.Fatal("Failed to connect to DB")
	}
	// init the functions
	env.Init()

	app := fiber.New()
	port := env.ENV.PORT

	routes.SetupRoutes(app)

	log.Printf("Geeye is running at %s\n", port)
	log.Fatal(app.Listen(":" + port))
}
