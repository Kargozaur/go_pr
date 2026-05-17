package svc

import (
	"context"
	"database/sql"
	"ecommerce/user-service/internal/jw"
	"ecommerce/user-service/internal/models"
	"ecommerce/user-service/internal/phasher"
	"ecommerce/user-service/internal/repo"
	"ecommerce/user-service/internal/repo/refresh"
	"ecommerce/user-service/internal/repo/user"
	"ecommerce/user-service/internal/schemas"
	"ecommerce/user-service/internal/util/validator"
	"errors"

	"github.com/uptrace/bun"
)

type UserService struct {
	db          *bun.DB
	userRepo    repo.IRepo
	profileRepo repo.IRepo
	refreshRepo repo.IRepo
	validator   validator.IValidator
	hasher      phasher.IHasher
	iss         jw.IJWT
}

func NewUserService(db *bun.DB, validator validator.IValidator, hasher phasher.IHasher, iss jw.IJWT) *UserService {
	userRepo := user.NewUserRepository(db)
	profileRepo := user.NewProfileRepo(db)
	refreshRepo := refresh.NewRefreshRepo(db)
	return &UserService{
		db:          db,
		userRepo:    userRepo,
		profileRepo: profileRepo,
		refreshRepo: refreshRepo,
		validator:   validator,
		hasher:      hasher,
		iss:         iss,
	}
}

func (u *UserService) Register(ctx context.Context, userData schemas.RegisterSchema) error {
	err := u.db.RunInTx(ctx, &sql.TxOptions{}, func(c context.Context, tx bun.Tx) error {
		err := u.validator.ValidateSchema(userData.UserDefaultSchema)
		if err != nil {
			return err
		}
		hashedPass, err := u.hasher.Hash(userData.Password)
		if err != nil {
			return err
		}
		userSchema := userData.ToUserSchema()
		userSchema.SwapWithHash(hashedPass)
		userModel, err := u.userRepo.Create(c, userSchema, tx)
		user, ok := userModel.(models.User)
		if !ok {
			return errors.New("Returned type does not match user model")
		}
		profileSchema := userData.ToUserProfileSchema()
		profileSchema.SetUUID(user.ID)
		_, err = u.profileRepo.Create(c, profileSchema, tx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
