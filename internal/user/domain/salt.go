package domain

import (
	"crypto/rand"
	"errors"
	"fmt"
)

const DefaultSizeSalt = 16

var ErrInvalidSalt = errors.New("invalid salt")

type Salt struct {
	value []byte
}

func GenerateRandomSalt(saltSize int) ([]byte, error) {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])

	if err != nil {
		return nil, err
	}

	return salt, nil
}

func NewSalt(value []byte) (Salt, error) {
	if len(value) != DefaultSizeSalt {
		return Salt{}, fmt.Errorf("%w: %s", ErrInvalidSalt, value)
	}

	return Salt{value}, nil
}

func (s Salt) ToPrimitive() []byte {
	return s.value
}
