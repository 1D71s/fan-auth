package services

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"myapp/database"
	"myapp/dto"
	"myapp/models"
	"myapp/utils"
	"strings"
	"time"
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

	email := strings.TrimSpace(data.Email)
	var existingUser models.User
	if err := database.DB.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "User with this email already exists",
		})
	}

	user := models.User{
		Email: email,
	}
	if err := user.SetPassword(data.Password); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to set password",
		})
	}
	err := database.DB.Create(&user)
	if err != nil {
		log.Println(err)
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Account created successfully!",
	})
}

func Login(c *fiber.Ctx) error {
	var data dto.LoginDto

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body:", err)
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	email := strings.TrimSpace(data.Email)
	var existingUser models.User
	if err := database.DB.Where("email = ?", email).First(&existingUser).Error; err != nil {
		fmt.Println("User not found or database error:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Email or Password",
		})
	}

	if err := existingUser.ComparePassword(strings.TrimSpace(data.Password)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Email or Password",
		})
	}

	token, err := utils.GenerateJWT(existingUser)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not generate token",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
	})
}
