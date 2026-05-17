package schemas

type RegisterSchema struct {
	UserDefaultSchema
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (r *RegisterSchema) ToUserSchema() *UserSchema {
	return &UserSchema{UserDefaultSchema: UserDefaultSchema{
		Email: r.Email, Password: r.Password,
	}}
}

func (r *RegisterSchema) ToUserProfileSchema() *Profile {
	return &Profile{FirstName: r.FirstName, LastName: r.LastName}
}
