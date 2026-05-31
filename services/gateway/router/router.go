package router

import (
	"net/http"
)

func RegisterUserRouter(mux *http.ServeMux) {
	mux.HandleFunc("POST /users/register", registerUser)
	mux.HandleFunc("POST /users/login", loginUser)
}
