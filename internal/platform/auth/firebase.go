package auth

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

func SetupFirebase() (*auth.Client, error) {
	// opt := option.WithCredentialsFile("./serviceAccountKey.json")

	app, err := firebase.NewApp(context.Background(), nil)

	if err != nil {
		return nil, err
	}

	return app.Auth(context.Background())
}
