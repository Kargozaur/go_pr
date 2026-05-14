package user

import (
	"context"
	m "ecommerce/user-service/internal/models"
	s "ecommerce/user-service/internal/schemas"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ProfileRepo struct {
	db *bun.DB
}

func NewProfileRepo(db *bun.DB) *ProfileRepo {
	return &ProfileRepo{db: db}
}

func (p *ProfileRepo) Create(ctx context.Context, profileSchema s.Profile, db bun.IDB) (*m.Profile, error) {
	profile := profileSchema.ToModel()
	err := db.NewInsert().
		Model(profile).
		Returning("*").
		Scan(ctx, profile)
	if err != nil {
		return nil, err
	}
	return profile, err
}

func (p *ProfileRepo) Read(ctx context.Context, userID uuid.UUID) (*m.Profile, error) {
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

func (p *ProfileRepo) Update(ctx context.Context, userID uuid.UUID, updateSchema s.UpdateProfile, db bun.IDB) (*m.Profile, error) {
	profile := updateSchema.ToModel()
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
func (p *ProfileRepo) Delete(ctx context.Context, userID uuid.UUID) (bool, error) {
	return true, nil
}
