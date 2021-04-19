package domain

import "context"

type UserRepository interface {
	Create(cxt context.Context, user *User) error
	FindByEmail(ctx context.Context, email Email) (*User, error)
}
