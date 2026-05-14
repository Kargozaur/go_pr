package schemas

import (
	"ecommerce/user-service/internal/models"

	"github.com/google/uuid"
)

type RefreshSchema struct {
	userID    uuid.UUID
	tokenHash string
}

func NewRefreshSchema(userID uuid.UUID, tokenHash string) RefreshSchema {
	return RefreshSchema{userID: userID, tokenHash: tokenHash}
}

func (r *RefreshSchema) ToModel() *models.RefreshTokens {
	return &models.RefreshTokens{
		UserID:    r.userID,
		TokenHash: r.tokenHash,
	}
}
