package domain

import (
	"errors"
	"fmt"
	"regexp"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

var ErrInvalidEmail = errors.New("invalid email")

type Email struct {
	value string
}

func NewEmail(value string) (Email, error) {
	isValid := isEmailValid(value)

	if !isValid {
		return Email{}, fmt.Errorf("%w: %s", ErrInvalidEmail, value)
	}

	return Email{value}, nil

}

func isEmailValid(email string) bool {
	if len(email) < 3 && len(email) > 150 {
		return false
	}

	return emailRegex.MatchString(email)
}

func (e Email) String() string {
	return e.value
}
