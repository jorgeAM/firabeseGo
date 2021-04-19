package domain

import (
	"crypto/sha512"
	"encoding/base64"
	"github/com/jorgeAM/goFireAuth/kit"
	"time"
)

type User struct {
	ID        kit.Id
	FirstName NoEmptyField
	LastName  NoEmptyField
	Email     Email
	Password  Password
	Salt      Salt
	CreatedAt time.Time
}

func NewUser(id, firstName, lastName, email, password string, salt []byte, createdAt time.Time) (*User, error) {
	idVO, err := kit.NewId(id)

	if err != nil {
		return nil, err
	}

	firstNameVO, err := NewNoEmptyField(firstName)

	if err != nil {
		return nil, err
	}

	lastNameVO, err := NewNoEmptyField(lastName)

	if err != nil {
		return nil, err
	}

	emailVO, err := NewEmail(email)

	if err != nil {
		return nil, err
	}

	passwordVO, err := NewPassword(password)

	if err != nil {
		return nil, err
	}

	saltVO, err := NewSalt(salt)

	if err != nil {
		return nil, err
	}

	return &User{
		ID:        idVO,
		FirstName: firstNameVO,
		LastName:  lastNameVO,
		Email:     emailVO,
		Password:  passwordVO,
		Salt:      saltVO,
		CreatedAt: createdAt,
	}, nil

}

func hashPassword(password string, salt []byte) string {
	passwordBytes := []byte(password)

	sha512Hasher := sha512.New()

	passwordBytes = append(passwordBytes, salt...)

	sha512Hasher.Write(passwordBytes)

	hashedPasswordBytes := sha512Hasher.Sum(nil)

	base64EncodedPasswordHash := base64.URLEncoding.EncodeToString(hashedPasswordBytes)

	return base64EncodedPasswordHash
}

func (u *User) SetHashPassword() (*User, error) {
	hashedPassword := hashPassword(u.Password.String(), u.Salt.ToPrimitive())

	hashedPasswordVO, err := NewPassword(hashedPassword)

	if err != nil {
		return nil, err
	}

	u.Password = hashedPasswordVO

	return u, nil
}

func (u *User) PasswordMatch(password string) bool {
	passwordHash := hashPassword(password, u.Salt.ToPrimitive())

	return u.Password.String() == passwordHash
}
