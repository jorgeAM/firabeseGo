package domain

import "context"

type UserRepository interface {
	Create(cxt context.Context, user *User) error
}
