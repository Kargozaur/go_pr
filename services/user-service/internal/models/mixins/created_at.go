package mixins

import "time"

type CreatedAt struct {
	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
}
