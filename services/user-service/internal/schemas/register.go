package schemas

type RegisterSchema struct {
	UserDefaultSchema
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}

func (r *RegisterSchema) ToUserSchema() *UserSchema {
	return &UserSchema{UserDefaultSchema: UserDefaultSchema{
		Email: r.Email, Password: r.Password,
	}}
}

func (r *RegisterSchema) ToUserProfileSchema() *Profile {
	return &Profile{FirstName: r.FirstName, LastName: r.LastName}
}
