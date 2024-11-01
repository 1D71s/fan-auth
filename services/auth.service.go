package services

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"myapp/dto"
)

func Register(c *fiber.Ctx) error {
	var data dto.RegisterDto

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body:", err)
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	if data.Password != data.RepeatPassword {
		return c.Status(400).JSON(fiber.Map{
			"error": "Passwords do not match",
		})
	}

	if len(data.Password) <= 6 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password must be greater than 6 characters",
		})
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Account created successfully!",
	})
}
