package main

import (
	"clean-arch/application"
	"clean-arch/domain"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize
	app := fiber.New()

	// Initialize application layer
	userService := application.NewUserService()

	// Define routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	app.Get("/users/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
		}

		user, err := userService.GetUser(id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString("User not found")
		}

		return c.JSON(user)
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		var user domain.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
		}

		if err := userService.CreateUser(&user); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Could not create user")
		}

		return c.Status(fiber.StatusCreated).JSON(user)
	})

	// Start the server
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
