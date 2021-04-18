package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type createTodoRequest struct {
	Title       string `json:"title" validate:"required,min=4,max=50"`
	Description string `json:"description,omitempty"`
}

func (h *Handler) CreateTodo(c *fiber.Ctx) error {
	var req createTodoRequest

	err := c.BodyParser(&req)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "something got wrong to parse request body"})
	}

	err = h.todoCreator.Create(c.Context(), req.Title, req.Description)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "todo created successfully"})
}

func (h *Handler) GetTodos(c *fiber.Ctx) error {
	todos, err := h.todoReader.Read(c.Context())

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"todos": todos})
}
