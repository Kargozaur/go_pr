package schemas

import (
	"ecommerce/user-service/internal/models"

	"github.com/google/uuid"
)

type Profile struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	userID    uuid.UUID `json:"-"`
}

type UpdateProfile struct {
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
}

func (p *Profile) SetUUID(userID uuid.UUID) {
	p.userID = userID
}

func (p *Profile) ToModel() any {
	if p.userID == uuid.Nil {
		return nil
	}
	return &models.Profile{
		FirstName: p.FirstName,
		LastName:  p.LastName,
		UserID:    p.userID,
	}
}

func (p *UpdateProfile) ToModel() any {
	profile := new(models.Profile)
	if p.FirstName != nil {
		profile.FirstName = *p.FirstName
	}
	if p.LastName != nil {
		profile.LastName = *p.LastName
	}
	return profile
}
