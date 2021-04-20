package application

import (
	"context"
	"errors"
	"github/com/jorgeAM/goFireAuth/internal/platform/auth"
	"github/com/jorgeAM/goFireAuth/internal/user/domain"
)

type UserFinder struct {
	repository domain.UserRepository
}

func NewUserFinder(repository domain.UserRepository) UserFinder {
	return UserFinder{repository}
}

func (u UserFinder) Find(ctx context.Context, email, password string) (*UserResponse, error) {
	emailVO, err := domain.NewEmail(email)

	if err != nil {
		return nil, err
	}

	user, err := u.repository.FindByEmail(ctx, emailVO)

	if err != nil {
		return nil, err
	}

	match := user.PasswordMatch(password)

	if !match {
		return nil, errors.New("Credentials are invalid")
	}

	jwt, err := auth.GenerateToken(user)

	if err != nil {
		return nil, err
	}

	userRes := newUserResponseFromAggregate(user, jwt)

	return userRes, nil
}
