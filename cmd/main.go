package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/gofiber/fiber/v2"
)

var (
	port = os.Getenv("PORT")
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":" + port)
}
