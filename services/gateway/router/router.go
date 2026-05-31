package router

import (
	"ecommerce/gateway/middleware"
	"net/http"
)

func RegisterUserRouter(mux *http.ServeMux) {
	mux.HandleFunc("POST /users/register", registerUser)
	mux.HandleFunc("POST /users/login", loginUser)
	mux.Handle("POST /users/logout/single", middleware.VerifyAccessToken(http.HandlerFunc(logoutUserSingle)))
	mux.Handle("POST /users/logout/all", middleware.VerifyAccessToken(http.HandlerFunc(logoutUserAll)))
}
