package jw

import (
	"github.com/google/uuid"
)

type IJWT interface {
	Issue(uuid.UUID, int64) (string, error)
	GetID(string) (uuid.UUID, error)
}
