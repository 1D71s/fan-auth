package services

import (
	"github.com/gofiber/fiber/v2"
)

func GetMe(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	email := c.Locals("email")

	c.Status(200)
	return c.JSON(fiber.Map{
		"user_id": userID,
		"email":   email,
		"message": "User information retrieved successfully",
	})
}
