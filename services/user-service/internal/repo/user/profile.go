package user

import (
	"context"
	m "ecommerce/user-service/internal/models"
	"ecommerce/user-service/internal/repo"
	"errors"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ProfileRepo struct {
	db *bun.DB
}

func NewProfileRepo(db *bun.DB) *ProfileRepo {
	return &ProfileRepo{db: db}
}

func (p *ProfileRepo) Create(ctx context.Context, profileSchema repo.RepoType, db bun.IDB) (any, error) {
	profile, ok := profileSchema.ToModel().(*m.Profile)
	if !ok {
		return nil, errors.New("Wrong model was passed")
	}
	err := db.NewInsert().
		Model(profile).
		Returning("*").
		Scan(ctx, profile)
	if err != nil {
		return nil, err
	}
	return profile, err
}

func (p *ProfileRepo) Read(ctx context.Context, userID any) (any, error) {
	if _, ok := userID.(uuid.UUID); !ok {
		return nil, errors.New("User ID must be passed as UUID")
	}
	profile := new(m.Profile)
	err := p.db.NewSelect().
		Model(profile).
		Where("user_id = ?", userID).
		Limit(1).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (p *ProfileRepo) Update(ctx context.Context, userID any, updateSchema repo.RepoType, db bun.IDB) (any, error) {
	if _, ok := userID.(uuid.UUID); !ok {
		return nil, errors.New("User ID must be passed as UUID")
	}
	profile, ok := updateSchema.ToModel().(*m.Profile)
	if !ok {
		return nil, errors.New("Wrong model was passed")
	}
	err := db.NewUpdate().
		Model(profile).
		Where("user_id = ?", userID).
		Returning("*").
		Scan(ctx, profile)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

// Implementation to satisfy the interface
func (p *ProfileRepo) Delete(ctx context.Context, userID any, db bun.IDB) (bool, error) {
	return true, nil
}
