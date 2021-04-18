package domain

import "context"

type TodoRepository interface {
	NewTodo(ctx context.Context, todo *Todo) error
	GetAllTodo(ctx context.Context) ([]*Todo, error)
}
