package auth

import (
	"context"

	_ "github.com/joho/godotenv/autoload"

	"github/com/jorgeAM/goFireAuth/internal/user/domain"
)

func GenerateToken(user *domain.User) (string, error) {
	authClient, err := SetupFirebase()

	if err != nil {
		return "", err
	}

	claims := map[string]interface{}{
		"displayName": user.FirstName.String() + user.LastName.String(),
	}

	token, err := authClient.CustomTokenWithClaims(context.Background(), user.ID.String(), claims)

	if err != nil {
		return "", err
	}

	return token, nil
}
