package schemas

import (
	"ecommerce/user-service/internal/models"

	"github.com/google/uuid"
)

// method SetUUID must be called before ToModel
type Profile struct {
	UserDefaultSchema
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	userID    uuid.UUID `json:"-"`
}

func (p *Profile) ToModel() *models.Profile {
	if p.userID == uuid.Nil {
		return nil
	}
	return &models.Profile{
		FirstName: p.FirstName,
		LastName:  p.LastName,
		UserID:    p.userID,
	}
}

func (p *Profile) SetUUID(userID uuid.UUID) {
	p.userID = userID
}
