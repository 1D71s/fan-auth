package routes

import (
	"github.com/gofiber/fiber/v2"
	"myapp/services"
)

func AuthRoutes(app *fiber.App) {
	app.Post("/api/auth/register", services.Register)
}
