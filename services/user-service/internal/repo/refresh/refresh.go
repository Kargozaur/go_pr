package refresh

import (
	"context"
	m "ecommerce/user-service/internal/models"
	s "ecommerce/user-service/internal/schemas"

	"github.com/uptrace/bun"
)

type RefreshRepo struct {
	db *bun.DB
}

func NewRefreshRepo(db *bun.DB) *RefreshRepo {
	return &RefreshRepo{db: db}
}

func (r *RefreshRepo) Create(ctx context.Context, refreshSchema s.RefreshSchema, db bun.IDB) (*m.RefreshTokens, error) {
	refreshToken := refreshSchema.ToModel()
	_, err := db.NewInsert().
		Model(refreshToken).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *RefreshRepo) Read(ctx context.Context, tokenHash string) (*m.RefreshTokens, error) {
	refreshToken := new(m.RefreshTokens)
	err := r.db.NewSelect().
		Model(refreshToken).
		Where("token_hash = ?", tokenHash).
		Scan(ctx, refreshToken)
	if err != nil {
		return nil, err
	}
	return refreshToken, nil
}

// Implementation ommited
func (r *RefreshRepo) Update(ctx context.Context, tokenHash string, schema s.RefreshSchema, db bun.IDB) (*m.RefreshTokens, error) {
	return nil, nil
}

func (r *RefreshRepo) Delete(ctx context.Context, tokenHash string, db bun.IDB) (bool, error) {
	refreshToken := new(m.RefreshTokens)
	result, err := db.NewDelete().
		Model(refreshToken).
		Where("token_hash = ?", tokenHash).
		Exec(ctx)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, nil
	}
	return affected == 1, nil
}
