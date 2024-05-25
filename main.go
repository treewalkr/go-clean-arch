package main

import (
	"clean-arch/adapters"
	"clean-arch/application"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize
	app := fiber.New()

	// Initialize application layer
	userService := application.NewUserService()
	userHandler := adapters.NewUserHandler(userService)

	// Define routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	app.Get("/users/:id", userHandler.GetUser)
	app.Post("/users", userHandler.CreateUser)

	// Start the server
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
