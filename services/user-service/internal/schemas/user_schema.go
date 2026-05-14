package schemas

import "ecommerce/user-service/internal/models"

type UserDefaultSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// method SwapWithHash must be called before ToModel()
// to ensure that the hashed password will be written in the database
type CreateUserSchema struct {
	UserDefaultSchema
}

func (c *CreateUserSchema) ToModel() *models.User {
	return &models.User{Email: c.Email, Password: c.Password}
}

func (c *CreateUserSchema) SwapWithHash(newPassword string) {
	c.Password = newPassword
}
