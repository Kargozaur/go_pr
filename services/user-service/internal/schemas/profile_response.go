package schemas

type UserData struct {
	Email string `json:"email"`
}

type ProfileResponse struct {
	Profile
	User UserData `json:"user"`
}
