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
	v, ok := userID.(uuid.UUID)
	if !ok {
		return nil, errors.New("User ID must be passed as UUID")
	}
	profile := new(m.Profile)
	err := p.db.NewSelect().
		Model(profile).
		Relation("User").
		Where("profiles.user_id = ?", v).
		Column("profiles.first_name", "profiles.last_name", "users.email").
		Limit(1).
		Scan(ctx, profile)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (p *ProfileRepo) Update(ctx context.Context, userID any, updateSchema repo.RepoType, db bun.IDB) (any, error) {
	v, ok := userID.(uuid.UUID)
	if !ok {
		return nil, errors.New("User ID must be passed as UUID")
	}
	profile, ok := updateSchema.ToModel().(*m.Profile)
	if !ok {
		return nil, errors.New("Wrong model was passed")
	}
	err := db.NewUpdate().
		Model(profile).
		Where("user_id = ?", v).
		Returning("*").
		Scan(ctx, profile)
	if err != nil {
		return nil, err
	}
	return profile, nil
}
