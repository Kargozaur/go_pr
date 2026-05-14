package schemas

import (
	"ecommerce/user-service/internal/models"

	"github.com/google/uuid"
)

type RefreshSchema struct {
	userID    uuid.UUID
	tokenHash string
}

func (r *RefreshSchema) ToModel() *models.RefreshTokens {
	return &models.RefreshTokens{
		UserID:    r.userID,
		TokenHash: r.tokenHash,
	}
}

func (r *RefreshSchema) SetUUID(userID uuid.UUID) {
	r.userID = userID
}

func (r *RefreshSchema) SetHash(tokenHash string) {
	r.tokenHash = tokenHash
}
