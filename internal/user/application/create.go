package application

import (
	"context"
	"github/com/jorgeAM/goFireAuth/internal/user/domain"
	"github/com/jorgeAM/goFireAuth/kit"
)

type UserCreator struct {
	repository domain.UserRepository
}

func NewUserCreator(repository domain.UserRepository) UserCreator {
	return UserCreator{repository}
}

func (u UserCreator) Create(ctx context.Context, firstName, lastName, email, password string) error {
	id := kit.NewID().String()

	user, err := domain.NewUser(id, firstName, lastName, email, password)

	if err != nil {
		return err
	}

	user, err = user.SetHashPassword()

	if err != nil {
		return err
	}

	return u.repository.Create(ctx, user)
}
