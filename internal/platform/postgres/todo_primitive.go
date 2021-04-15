package postgres

import (
	"github/com/jorgeAM/goFireAuth/internal/todo"
	"time"
)

type Todo struct {
	id          string
	title       string
	description string
	createdAt   time.Time
}

func (t Todo) UnmarshalAggregate() (*todo.Todo, error) {
	todo, err := todo.NewTodo(t.id, t.title, t.description)

	if err != nil {
		return nil, err
	}

	return todo, nil
}
