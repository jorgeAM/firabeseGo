package kit

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var ErrInvalidID = errors.New("invalid user ID")

type Id struct {
	value string
}

func NewId(value string) (Id, error) {
	v, err := uuid.Parse(value)

	if err != nil {
		return Id{}, fmt.Errorf("%w: %s", ErrInvalidID, value)
	}

	return Id{value: v.String()}, nil
}

func NewID() Id {
	v := uuid.New()
	return Id{value: v.String()}
}

func (i Id) String() string {
	return i.value
}
