package models

import (
	"ecommerce/user-service/internal/models/mixins"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`
	Email         string `bun:"email,unique"`
	Password      string `bun:"password,notnull"`
	mixins.IDMixin
	mixins.UpdatedAt
	mixins.CreatedAt

	Profile       *Profile       `bun:"rel:has-one,join:id=user_id"`
	RefreshTokens *RefreshTokens `bun:"rel:has-many,join:id=user_id"`
}

type Profile struct {
	bun.BaseModel `bun:"table:profiles"`
	FirstName     string    `bun:"first_name,notnull,type:varchar(255)"`
	LastName      string    `bun:"last_name,notnull,type:varchar(255)"`
	UserID        uuid.UUID `bun:"user_id,unique"`
	mixins.UpdatedAt
	mixins.IDMixin
}
