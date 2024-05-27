package main

import (
	"clean-arch/adapters"
	"clean-arch/application"
	"clean-arch/domain"
	"clean-arch/infrastructure/persistence"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Initialize the Infrastructure Layer (In-Memory DB)
	var userRepo domain.UserRepository = persistence.NewInMemoryUserRepository()

	// Initialize the Application Layer
	userService := application.NewUserService(userRepo)

	// Initialize the Adapters (or interface) Layer
	userHandler := adapters.NewUserHandler(userService)

	// Set up the routes and assign handlers
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
