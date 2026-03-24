package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ZuhybDev/geeyeApp/database"
	"github.com/ZuhybDev/geeyeApp/handlers"
	"github.com/gofiber/fiber/v3"
)

func handlerFunc(c fiber.Ctx) error {
	return c.SendString("Hello, from handle func")

}

func main() {

	// ENV
	// err := godotenv.Load()

	// if err != nil {
	// 	log.Fatal("Unable to load ENV environments")
	// }

	// Connect once!
	database.Connect()

	app := fiber.New()
	port := os.Getenv("PORT")

	fmt.Println("DBURL from env is:", os.Getenv("DATABASE_URL"))

	// routes
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World")
	})

	app.Get("/api/hello", handlers.HelloRoute)

	log.Fatal(app.Listen(port))
}
