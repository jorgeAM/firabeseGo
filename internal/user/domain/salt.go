package domain

import (
	"crypto/rand"
)

const saltSize = 16

type Salt struct {
	value []byte
}

func generateRandomSalt(saltSize int) ([]byte, error) {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])

	if err != nil {
		return nil, err
	}

	return salt, nil
}

func NewSalt() (Salt, error) {
	saltBytes, err := generateRandomSalt(saltSize)

	if err != nil {
		return Salt{}, err
	}

	return Salt{value: saltBytes}, nil
}

func (s Salt) ToPrimitive() []byte {
	return s.value
}
