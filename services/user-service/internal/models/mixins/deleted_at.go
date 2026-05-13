package mixins

import (
	"database/sql"
)

type DeletedAtMixin struct {
	DeletedAT sql.NullTime `bun:"deleted_at"`
}
