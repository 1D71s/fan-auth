package routes

import (
	"github.com/gofiber/fiber/v2"
	"myapp/middlewares"
	"myapp/services"
)

func AuthRoutes(app *fiber.App) {
	app.Post("/api/auth/register", services.Register)
	app.Post("/api/auth/login", services.Login)

	protected := app.Group("/protected", middlewares.JWTAuth)

	protected.Get("api/users/me", services.GetMe)
}
