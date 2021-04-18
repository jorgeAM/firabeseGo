package domain

import (
	"errors"
	"fmt"
)

var ErrInvalidPassword = errors.New("invalid password")

type Password struct {
	value string
}

func NewPassword(value string) (Password, error) {
	if len(value) < 6 {
		return Password{}, fmt.Errorf("%w: %s", ErrInvalidPassword, value)
	}

	return Password{value}, nil
}

func (p Password) String() string {
	return p.value
}
