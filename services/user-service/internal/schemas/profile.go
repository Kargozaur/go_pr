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

func NewProfile(firstName, lastName string, userID uuid.UUID) Profile {
	return Profile{FirstName: firstName, LastName: lastName, userID: userID}
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

func (p *UpdateProfile) ToModel() *models.Profile {
	profile := new(models.Profile)
	if p.FirstName != nil {
		profile.FirstName = *p.FirstName
	}
	if p.LastName != nil {
		profile.LastName = *p.LastName
	}
	return profile
}
