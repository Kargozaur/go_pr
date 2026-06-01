package refresh

import (
	"context"
	m "ecommerce/user-service/internal/models"
	"ecommerce/user-service/internal/repo"
	"errors"

	"github.com/google/uuid"
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

func (r *RefreshRepo) Delete(ctx context.Context, identifier any, db bun.IDB) (bool, error) {
	query := db.NewDelete()
	switch v := identifier.(type) {
	case uuid.UUID:
		query = query.Where("user_id = ?", v)
	case string:
		query = query.Where("token_hash = ?", v)
	default:
		return false, errors.New("Either user id or token hash must be passed")
	}
	result, err := query.Exec(ctx)
	if err != nil {
		return false, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if _, ok := identifier.(string); ok {
		return affected == 1, nil
	} else {
		return affected >= 1, nil
	}
}
