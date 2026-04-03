package main

import (
	"log"

	connection "github.com/ZuhybDev/geeyeApp/config"
	env "github.com/ZuhybDev/geeyeApp/envConfig"
	"github.com/ZuhybDev/geeyeApp/routes"
	"github.com/gofiber/fiber/v3"
)

func main() {

	// init the functions
	env.Init()

	// Connect once
	connection.Connect()
	if connection.DBPool == nil {
		log.Fatal("Failed to connect to DB")
	}

	app := fiber.New()
	port := env.ENV.PORT

	routes.SetupRoutes(app)

	log.Printf("Geeye is running at %s\n", port)
	log.Fatal(app.Listen(":" + port))
}
