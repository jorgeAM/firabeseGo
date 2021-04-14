package todo

import (
	"github/com/jorgeAM/goFireAuth/kit"
	"time"
)

type Todo struct {
	ID          kit.Id
	Title       string
	Description string
	CreatedAt   time.Time
}

func NewTodo(id, title, description string) (*Todo, error) {
	idVO, err := kit.NewId(id)

	if err != nil {
		return nil, err
	}

	return &Todo{
		ID:          idVO,
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
	}, nil

}
