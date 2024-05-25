package interfaces

import (
    "clean-arch/application"
    "clean-arch/domain"
    "github.com/gofiber/fiber/v2"
    "strconv"
)

type UserHandler struct {
    userService application.UserService
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(userService application.UserService) *UserHandler {
    return &UserHandler{
        userService: userService,
    }
}

// GetUser handles GET requests to retrieve a user by ID
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
    }

    user, err := h.userService.GetUser(id)
    if err != nil {
        return c.Status(fiber.StatusNotFound).SendString("User not found")
    }

    return c.JSON(user)
}

// CreateUser handles POST requests to create a new user
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
    var user domain.User
    if err := c.BodyParser(&user); err != nil {
        return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
    }

    if err := h.userService.CreateUser(&user); err != nil {
        return c.Status(fiber.StatusInternalServerError).SendString("Could not create user")
    }

    return c.Status(fiber.StatusCreated).JSON(user)
}
