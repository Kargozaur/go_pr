package mixins

import "time"

type UpdatedAt struct {
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}
