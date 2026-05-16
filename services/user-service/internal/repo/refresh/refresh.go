package refresh

import (
	"context"
	m "ecommerce/user-service/internal/models"
	"ecommerce/user-service/internal/repo"
	"errors"

	"github.com/uptrace/bun"
)

type RefreshRepo struct {
	db *bun.DB
}

func NewRefreshRepo(db *bun.DB) *RefreshRepo {
	return &RefreshRepo{db: db}
}

func (r *RefreshRepo) Create(ctx context.Context, refreshSchema repo.RepoType, db bun.IDB) (any, error) {
	refreshToken, ok := refreshSchema.ToModel().(m.RefreshTokens)
	if !ok {
		return nil, errors.New("Wrong model was passed")
	}
	_, err := db.NewInsert().
		Model(refreshToken).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *RefreshRepo) Read(ctx context.Context, tokenHash any) (any, error) {
	if _, ok := tokenHash.(string); !ok {
		return nil, errors.New("token hash must be passed as string")
	}
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
func (r *RefreshRepo) Update(ctx context.Context, tokenHash any, schema repo.RepoType, db bun.IDB) (any, error) {
	return nil, nil
}

func (r *RefreshRepo) Delete(ctx context.Context, tokenHash any, db bun.IDB) (bool, error) {
	if _, ok := tokenHash.(string); !ok {
		return false, errors.New("Token hash must be passed as string")
	}
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
