package auth

import (
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"github/com/jorgeAM/goFireAuth/internal/user/domain"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = os.Getenv("JWT_SECRET")

func GenerateToken(user *domain.User) (string, error) {
	displayName := user.FirstName.String() + user.LastName.String()

	claim := Claim{
		ID:          user.ID.String(),
		DisplayName: displayName,
		Email:       user.Email.String(),
		StandardClaims: jwt.StandardClaims{
			Issuer:    "go-firebase",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(8 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	jwt, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", err
	}

	return jwt, nil
}
