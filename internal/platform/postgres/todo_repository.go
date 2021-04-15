package postgres

import (
	"context"
	"fmt"
	"github/com/jorgeAM/goFireAuth/internal/todo/domain"

	"github.com/go-pg/pg/v10"
)

type todoRepository struct {
	db *pg.DB
}

func NewRepository(db *pg.DB) domain.Repository {
	return &todoRepository{db}
}

func (r *todoRepository) NewTodo(ctx context.Context, todo *domain.Todo) error {
	_, err := r.db.ModelContext(ctx, &Todo{
		ID:          todo.ID.String(),
		Title:       todo.Title,
		Description: todo.Description,
		CreatedAt:   todo.CreatedAt,
	}).Insert()

	if err != nil {
		return err
	}

	fmt.Println("si todo va bien esto deberias verlo")
	return nil
}

func (r *todoRepository) GetAllTodo(ctx context.Context) ([]*domain.Todo, error) {
	var todos []Todo

	err := r.db.ModelContext(ctx, &todos).Select()

	if err != nil {
		return nil, err
	}

	var todoEntities []*domain.Todo

	for _, t := range todos {
		todoEntity, err := domain.NewTodo(t.ID, t.Title, t.Description, t.CreatedAt)

		if err != nil {
			return nil, err
		}

		todoEntities = append(todoEntities, todoEntity)
	}

	return todoEntities, nil
}
