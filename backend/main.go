package main

import (
	"backend/database"
	"backend/internal/router"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	app := fiber.New()
	database.Init()
	router.InitRouter(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to FastTea API")
	})
	app.Listen(":8080")
}
