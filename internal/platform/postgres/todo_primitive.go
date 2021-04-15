package postgres

import (
	"github/com/jorgeAM/goFireAuth/internal/todo/domain"
	"time"
)

type Todo struct {
	id          string
	title       string
	description string
	createdAt   time.Time
}

func (t Todo) UnmarshalAggregate() (*domain.Todo, error) {
	todo, err := domain.NewTodo(t.id, t.title, t.description, t.createdAt)

	if err != nil {
		return nil, err
	}

	return todo, nil
}
