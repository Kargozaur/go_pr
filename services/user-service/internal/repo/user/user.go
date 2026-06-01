package user

import (
	"context"
	"ecommerce/user-service/internal/models"
	m "ecommerce/user-service/internal/models"
	"ecommerce/user-service/internal/repo"
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

func (u *UserRepository) Create(ctx context.Context, userSchema repo.RepoType, db bun.IDB) (any, error) {
	user, ok := userSchema.ToModel().(*models.User)
	if !ok {
		return nil, errors.New("Wrong model passed")
	}
	err := db.NewInsert().Model(user).Returning("*").Scan(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) Read(ctx context.Context, email any) (any, error) {
	if _, ok := email.(string); !ok {
		return nil, errors.New("email must be passed as a string")
	}
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

func (u *UserRepository) Update(ctx context.Context, email any, userSchema repo.RepoType, db bun.IDB) (any, error) {
	if _, ok := email.(string); !ok {
		return nil, errors.New("email must be passed as a string")
	}
	user, ok := userSchema.ToModel().(*models.User)
	if !ok {
		return nil, errors.New("Wrong model passed")
	}
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
