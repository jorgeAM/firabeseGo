package application

import (
	"github/com/jorgeAM/goFireAuth/internal/todo/domain"
	"time"
)

type TodoResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

func newTodoResponseFromAggregate(todo *domain.Todo) *TodoResponse {
	return &TodoResponse{
		ID:          todo.ID.String(),
		Title:       todo.Title,
		Description: todo.Description,
		CreatedAt:   todo.CreatedAt,
	}
}
