package schemas

import "ecommerce/user-service/internal/models"

type UserDefaultSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// method SwapWithHash must be called before ToModel()
// to ensure that the hashed password will be written in the database
type UserSchema struct {
	UserDefaultSchema
}

type UserUpdateSchema struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

func (c *UserSchema) ToModel() any {
	return &models.User{Email: c.Email, Password: c.Password}
}

func (c *UserSchema) SwapWithHash(newPassword string) {
	c.Password = newPassword
}

func (c *UserUpdateSchema) ToModel() any {
	user := new(models.User)
	if c.Email != nil {
		user.Email = *c.Email
	}
	if c.Password != nil {
		user.Password = *c.Password
	}
	return user
}

func (c *UserUpdateSchema) SwapPassword(passwordHash string) {
	if c.Password != nil {
		c.Password = &passwordHash
	}
}
