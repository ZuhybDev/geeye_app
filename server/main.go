package main

import (
	"log"
	"os"

	connection "github.com/ZuhybDev/geeyeApp/config"
	"github.com/ZuhybDev/geeyeApp/handlers"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {

	// ENV
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Unable to load ENV environments")
	}

	// Connect once!
	connection.Connect()

	if connection.DBPool == nil {
		log.Fatal("Failed to connect to DB")
	}

	app := fiber.New()
	port := os.Getenv("PORT")

	// fmt.Println("DBURL from env is:", os.Getenv("DATABASE_URL"))

	// routes
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World")
	})

	app.Get("/api/all-users", handlers.GetListUsers)

	log.Fatal(app.Listen(":" + port))

	log.Printf("Geeye is running at %s\n", port)
}
