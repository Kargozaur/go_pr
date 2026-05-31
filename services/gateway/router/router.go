package router

import (
	"net/http"
)

func RegisterUserRouter(mux *http.ServeMux) {
	mux.HandleFunc("POST /users/register", registerUser)
	mux.HandleFunc("POST /users/login", loginUser)
	mux.HandleFunc("POST /users/logout/single", logoutUserSingle)
	mux.HandleFunc("POST /users/logout/all", logoutUserAll)
}
