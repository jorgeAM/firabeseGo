package postgres

import (
	"context"
	"github/com/jorgeAM/goFireAuth/internal/user/domain"

	"github.com/go-pg/pg/v10"
)

type userRepository struct {
	db *pg.DB
}

func NewUserRepository(db *pg.DB) domain.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	_, err := r.db.ModelContext(ctx, &User{
		ID:        user.ID.String(),
		FirstName: user.FirstName.String(),
		LastName:  user.LastName.String(),
		Email:     user.Email.String(),
		Password:  user.Password.String(),
		Salt:      user.Salt.ToPrimitive(),
		CreatedAt: user.CreatedAt,
	}).Insert()

	if err != nil {
		return err
	}

	return nil
}
