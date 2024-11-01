package main

import (
	"github.com/joho/godotenv"
	"log"
	"myapp/database"
	"myapp/routes"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn`t read .evn file")
	}
	port := os.Getenv("PORT")
	app := fiber.New()

	routes.AuthRoutes(app)

	log.Fatal(app.Listen(":" + port))
}
