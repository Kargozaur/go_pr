package jw_test

import (
	"ecommerce/user-service/internal/jw"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func TestJWT(t *testing.T) {
	os.Setenv("SECRET_KEY", "secret-key")
	defer os.Unsetenv("SECRET_KEY")
	jwtProvider := jw.NewJWT()
	newUID, err := uuid.NewV7()
	if err != nil {
		t.Fatal(err)
	}
	t.Run("CreateAndVerify", func(t *testing.T) {
		tokenString, err := jwtProvider.Issue(newUID, time.Now().Add(time.Hour).UTC().Unix())
		if err != nil {
			t.Fatal(err)
		}
		if _, err := jwtProvider.GetID(tokenString); err != nil {
			t.Fatal("Failed to verify access token")
		}
		uId, err := jwtProvider.GetID(tokenString)
		if err != nil {
			t.Fatal(err)
		}
		if uId != newUID {
			t.Fatal("Failed to get correct user ID from token")
		}
	})
	t.Run("InvalidToken", func(t *testing.T) {
		if _, err := jwtProvider.GetID("invalid_token"); err == nil {
			t.Fatal("Expected invalid token to fail verification")
		}
	})
	t.Run("CheckExpired", func(t *testing.T) {
		expiredClaims := jwt.MapClaims{
			"sub": newUID.String(),
			"exp": time.Now().Add(-1 * time.Hour).UTC().Unix(),
		}
		secret := []byte(os.Getenv("SECRET_KEY"))
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, expiredClaims)
		tokenString, err := token.SignedString(secret)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := jwtProvider.GetID(tokenString); err == nil {
			t.Fatal("Expected expired token to fail verification")
		}
	})
}
