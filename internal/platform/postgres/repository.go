package postgres

import (
	"context"
	"github/com/jorgeAM/goFireAuth/internal/todo"

	"github.com/go-pg/pg/v10"
)

type repository struct {
	db *pg.DB
}

func NewRepository(db *pg.DB) todo.Repository {
	return &repository{db}
}

func (r *repository) NewTodo(ctx context.Context, todo *todo.Todo) error {
	return nil
}

func (r *repository) GetAllTodo(ctx context.Context) ([]*todo.Todo, error) {
	return nil, nil
}
