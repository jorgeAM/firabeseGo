package postgres

import (
	"github/com/jorgeAM/goFireAuth/internal/user/domain"
	"time"
)

const saltSize = 16

type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string
	Salt      []byte
	CreatedAt time.Time
}

func (u User) UnmarshalAggregate() (*domain.User, error) {
	user, err := domain.NewUser(u.ID, u.FirstName, u.LastName, u.Email, u.Password, u.Salt, u.CreatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}
