package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type signUpRequest struct {
	FirstName string `json:"firstName" validate:"required,min=2,max=50"`
	LastName  string `json:"lastName" validate:"required,min=2,max=50"`
	Email     string `json:"email" validate:"required,email,min=6"`
	Password  string `json:"password" validate:"required,min=6"`
}

func (h *Handler) SignUp(c *fiber.Ctx) error {
	var req signUpRequest

	err := c.BodyParser(&req)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "something got wrong to parse request body"})
	}

	fmt.Println("should not print")

	err = h.userCreator.Create(c.Context(), req.FirstName, req.LastName, req.Email, req.Password)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "user created successfully"})
}
