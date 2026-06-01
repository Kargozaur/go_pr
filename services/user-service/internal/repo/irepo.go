package repo

import (
	"context"

	"github.com/uptrace/bun"
)

type Create interface {
	Create(context.Context, RepoType, bun.IDB) (any, error)
}

type Read interface {
	Read(context.Context, any) (any, error)
}

type Update interface {
	Update(context.Context, any, RepoType, bun.IDB) (any, error)
}

type Delete interface {
	Delete(context.Context, any, bun.IDB) (bool, error)
}

type CRU interface {
	Create
	Read
	Update
}

type CRD interface {
	Create
	Read
	Delete
}
