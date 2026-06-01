package router

import (
	"ecommerce/gateway/middleware"
	"net/http"
)

func RegisterUserRouter(mux *http.ServeMux) {
	handler := NewHandler()
	mux.HandleFunc("POST /users/register", handler.registerUser)
	mux.HandleFunc("POST /users/login", handler.loginUser)
	mux.Handle("POST /users/logout/single", middleware.VerifyAccessToken(http.HandlerFunc(handler.logoutUserSingle)))
	mux.Handle("POST /users/logout/all", middleware.VerifyAccessToken(http.HandlerFunc(handler.logoutUserAll)))
	mux.Handle("GET /users/profile/me", middleware.VerifyAccessToken(http.HandlerFunc(handler.GetProfile)))
}
