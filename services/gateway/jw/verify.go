package jw

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type Verifier struct {
	Secret string
}

func NewVerifier() *Verifier {
	godotenv.Load(".env")
	return &Verifier{
		Secret: os.Getenv("SECRET_KEY"),
	}
}

func (v *Verifier) Verify(token string) (bool, error) {
	_, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		return []byte(v.Secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
