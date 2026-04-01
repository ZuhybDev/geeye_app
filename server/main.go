package main

import (
	"log"
	"os"

	connection "github.com/ZuhybDev/geeyeApp/config"
	"github.com/ZuhybDev/geeyeApp/routes"
	"github.com/ZuhybDev/geeyeApp/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {

	// ENV
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load ENV environments")
	}

	utils.LoadConfig()

	// Connect once!
	connection.Connect()
	if connection.DBPool == nil {
		log.Fatal("Failed to connect to DB")
	}

	app := fiber.New()
	port := os.Getenv("PORT")

	routes.SetupRoutes(app)

	log.Printf("Geeye is running at %s\n", port)
	log.Fatal(app.Listen(":" + port))
}
