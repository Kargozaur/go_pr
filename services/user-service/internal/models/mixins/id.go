package mixins

import "github.com/google/uuid"

type IDMixin struct {
	ID uuid.UUID `bun:"id,pk,default:uuidv7"`
}
