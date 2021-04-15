package postgres

import (
	"github/com/jorgeAM/goFireAuth/internal/todo/domain"
	"time"
)

type Todo struct {
	ID          string
	Title       string
	Description string
	CreatedAt   time.Time
}

func (t Todo) UnmarshalAggregate() (*domain.Todo, error) {
	todo, err := domain.NewTodo(t.ID, t.Title, t.Description, t.CreatedAt)

	if err != nil {
		return nil, err
	}

	return todo, nil
}
