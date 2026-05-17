package thasher

import (
	"crypto/sha256"
	"encoding/hex"
)

type Hasher struct{}

func NewHasher() *Hasher {
	return &Hasher{}
}

func (h *Hasher) Hash(tokenString string) string {
	data := []byte(tokenString)
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}
