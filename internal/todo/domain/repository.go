package domain

import "context"

type Repository interface {
	NewTodo(ctx context.Context, todo *Todo) error
	GetAllTodo(ctx context.Context) ([]*Todo, error)
}
