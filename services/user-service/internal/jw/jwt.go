package jw

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWT struct {
	SecretKey []byte
}

func NewJWT() *JWT {
	return &JWT{SecretKey: []byte(os.Getenv("JWT_KEY"))}
}

func (j *JWT) Issue(userID uuid.UUID, t int) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID.String(),
		"iat": time.Now().UTC().Unix(),
		"exp": t,
		"iss": "user-service",
	})
	return claims.SignedString(j.SecretKey)
}

func (j *JWT) GetID(tokenString string) (uuid.UUID, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return j.SecretKey, nil
	})
	if err != nil {
		return uuid.Nil, err
	}
	if !token.Valid {
		return uuid.Nil, errors.New("Invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.Nil, jwt.ErrInvalidType
	}
	sub, ok := claims["sub"].(string)
	if !ok {
		return uuid.Nil, jwt.ErrInvalidType
	}
	return uuid.Parse(sub)
}
