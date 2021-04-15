package application

import (
	"context"
	"github/com/jorgeAM/goFireAuth/internal/todo/domain"
)

type TodoReader struct {
	repository domain.Repository
}

func NewTodoReader(repository domain.Repository) TodoReader {
	return TodoReader{repository}
}

func (t TodoReader) Read(ctx context.Context) ([]*TodoResponse, error) {
	todosAggregate, err := t.repository.GetAllTodo(ctx)

	if err != nil {
		return nil, err
	}

	var todosResponse []*TodoResponse

	for _, aggregate := range todosAggregate {
		todoResponse := newTodoResponseFromAggregate(aggregate)
		todosResponse = append(todosResponse, todoResponse)
	}

	return todosResponse, nil
}
