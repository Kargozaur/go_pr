package models

import (
	"ecommerce/user-service/internal/models/mixins"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RefreshTokens struct {
	bun.BaseModel
	TokenHash string    `bun:"notnull,type:varchar(150)"`
	UserID    uuid.UUID `bun:"user_id,notnull"`
	mixins.IDMixin
	mixins.CreatedAt
	mixins.DeletedAtMixin
}
