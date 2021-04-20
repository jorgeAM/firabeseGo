package auth

import "github.com/dgrijalva/jwt-go"

type Claim struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	jwt.StandardClaims
}
