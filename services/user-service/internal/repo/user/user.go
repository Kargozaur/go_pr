package user

import (
	"context"
	m "ecommerce/user-service/internal/models"
	s "ecommerce/user-service/internal/schemas"
	"errors"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type UserRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) Create(ctx context.Context, userSchema s.UserSchema, db bun.IDB) (*m.User, error) {
	user := userSchema.ToModel()
	err := db.NewInsert().Model(user).Returning("*").Scan(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) Read(ctx context.Context, email string) (*m.User, error) {
	user := new(m.User)
	err := u.db.NewSelect().
		Model(user).
		Where("email = ?", email).
		Limit(1).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) Update(ctx context.Context, email string, userSchema s.UserUpdateSchema, db bun.IDB) (*m.User, error) {
	user := userSchema.ToModel()
	err := db.NewUpdate().
		Model(user).
		Where("email = ?", email).
		Returning("*").
		Scan(ctx, user)
	if err != nil {
		return nil, err
	}
	if user.ID == uuid.Nil {
		return nil, errors.New("User not found")
	}
	return user, nil
}

// Method to satisfy the interface requirements. No implementation is provided
func (u *UserRepository) Delete(ctx context.Context, userID uuid.UUID, db bun.IDB) (bool, error) {
	return true, nil
}
