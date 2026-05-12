package phasher_test

import (
	"ecommerce/user-service/internal/phasher"
	"testing"
)

func TestHasher(t *testing.T) {
	hasher := phasher.NewPasswordHasher(12)
	tests := []struct {
		name     string
		password string
	}{
		{name: "test1", password: "password1"},
		{name: "test2", password: "password2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash, err := hasher.Hash(tt.password)
			if err != nil {
				t.Fatal(err)
			}
			if ok := hasher.VerifyPassword(tt.password, hash); ok != true {
				t.Fatal(ok)
			}
		})
	}
}
