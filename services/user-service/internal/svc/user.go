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
	"ecommerce/user-service/internal/thasher"
	"ecommerce/user-service/internal/util/validator"
	"errors"
	"time"

	"github.com/uptrace/bun"
)

type UserService struct {
	db          *bun.DB
	userRepo    repo.IRepo
	profileRepo repo.IRepo
	refreshRepo repo.IRepo
	validator   validator.IValidator
	hasher      phasher.IHasher
	thasher     thasher.IHasher
	iss         jw.IJWT
}

func NewUserService(db *bun.DB, validator validator.IValidator, hasher phasher.IHasher,
	thasher thasher.IHasher, iss jw.IJWT) *UserService {
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
		thasher:     thasher,
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
		if err != nil {
			return err
		}
		user, ok := userModel.(models.User)
		if !ok {
			return errors.New("Returned type doesn't match the user model")
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

func (u *UserService) Login(ctx context.Context, loginSchema schemas.LoginSchema) (*schemas.TokenResponse, error) {
	userModel, err := u.userRepo.Read(ctx, loginSchema.Email)
	if err != nil {
		return nil, errors.New("Invalid credentials")
	}
	user, ok := userModel.(models.User)
	if !ok {
		return nil, errors.New("Returned type doesn't match the user model")
	}
	if ok := u.hasher.VerifyPassword(loginSchema.Password, user.Password); !ok {
		return nil, errors.New("Invalid credentials")
	}
	if err != nil {
		return nil, err
	}
	accessToken, err := u.iss.Issue(user.ID, time.Now().UTC().Add(time.Minute*30).Unix())
	if err != nil {
		return nil, err
	}
	refreshToken, err := u.iss.Issue(user.ID, time.Now().UTC().Add(time.Hour*72).Unix())
	if err != nil {
		return nil, err
	}
	refreshSchema := schemas.NewRefreshSchema(user.ID, u.thasher.Hash(refreshToken))
	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{})
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	_, err = u.refreshRepo.Create(ctx, refreshSchema, tx)
	if err != nil {
		return nil, err
	}
	return &schemas.TokenResponse{AccessToken: accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
	}, nil
}

func (u *UserService) Logout(ctx context.Context, identifier any) error {
	tx, err := u.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	res, err := u.refreshRepo.Delete(ctx, identifier, tx)
	if err != nil {
		return err
	}
	if !res {
		return errors.New("Failed to logout the user")
	}
	return nil
}
