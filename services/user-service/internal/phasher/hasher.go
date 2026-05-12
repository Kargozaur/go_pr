package phasher

import "golang.org/x/crypto/bcrypt"

type Hasher struct {
	cost int
}

func NewPasswordHasher(cost int) *Hasher {
	return &Hasher{cost: cost}
}

func (h *Hasher) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (h *Hasher) VerifyPassword(password, passwordHash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)) == nil
}
