package main

import (
	"fmt"
	"log"
	"request_limiter/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	app := fiber.New()

	if app != nil {
		fmt.Println("Fiber server is running on port 5000")
		app.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("Hello, World !")
		})

		routes.ItemRoutes(app)

		app.Listen(":5000")
	}

}
