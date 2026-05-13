package jw

import (
	"github.com/google/uuid"
)

type IJWT interface {
	Issue(uuid.UUID, int) (string, error)
	GetID(string) (uuid.UUID, error)
}
