package repo

import (
	"context"

	"github.com/uptrace/bun"
)

// General interface to describe the repository. It may be separated into 4 different
// interfaces to follow the Interface segregation principle in SOLID.
type IRepo interface {
	Create(context.Context, RepoType, bun.IDB) (any, error)
	Read(context.Context, any) (any, error)
	Update(context.Context, any, RepoType, bun.IDB) (any, error)
	Delete(context.Context, any, bun.IDB) (bool, error)
}
