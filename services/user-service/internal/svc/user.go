package svc

import (
	"ecommerce/user-service/internal/jw"
	"ecommerce/user-service/internal/phasher"
	"ecommerce/user-service/internal/repo"
	"ecommerce/user-service/internal/repo/refresh"
	"ecommerce/user-service/internal/repo/user"
	"ecommerce/user-service/internal/util/validator"

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
