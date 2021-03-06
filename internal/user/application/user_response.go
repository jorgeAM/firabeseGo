package application

import (
	"github/com/jorgeAM/goFireAuth/internal/user/domain"
	"time"
)

type UserResponse struct {
	ID        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"LastName"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"createdAt"`
}

func newUserResponseFromAggregate(user *domain.User, jwt string) *UserResponse {
	return &UserResponse{
		ID:        user.ID.String(),
		FirstName: user.FirstName.String(),
		LastName:  user.LastName.String(),
		Email:     user.Email.String(),
		Token:     jwt,
		CreatedAt: user.CreatedAt,
	}
}
