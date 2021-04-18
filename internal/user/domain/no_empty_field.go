package domain

import (
	"errors"
	"fmt"
)

var ErrInvalidField = errors.New("invalid field")

type NoEmptyField struct {
	value string
}

func NewNoEmptyField(value string) (NoEmptyField, error) {
	if len(value) == 0 {
		return NoEmptyField{}, fmt.Errorf("%w: %s", ErrInvalidField, value)
	}

	return NoEmptyField{value}, nil

}

func (e NoEmptyField) String() string {
	return e.value
}
