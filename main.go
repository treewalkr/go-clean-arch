package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize
	app := fiber.New()

	// Define a route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	// Start the server
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
