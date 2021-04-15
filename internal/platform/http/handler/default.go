package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) SayHi(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
