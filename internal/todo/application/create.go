package application

import (
	"context"
	"time"

	"github/com/jorgeAM/goFireAuth/internal/todo/domain"
	"github/com/jorgeAM/goFireAuth/kit"
)

type TodoCreator struct {
	repository domain.TodoRepository
}

func NewTodoCreator(repository domain.TodoRepository) TodoCreator {
	return TodoCreator{repository}
}

func (t TodoCreator) Create(ctx context.Context, title, description string) error {
	id := kit.NewID().String()

	todo, err := domain.NewTodo(id, title, description, time.Now())

	if err != nil {
		return err
	}

	return t.repository.NewTodo(ctx, todo)
}
