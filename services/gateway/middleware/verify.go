package middleware

import "net/http"

func VerifyAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}
		cookie, err := r.Cookie("access_token")
		if err != nil {
			http.Error(w, "Access token cookie is required", http.StatusUnauthorized)
			return
		}
		if cookie.Value != token || cookie == nil || cookie.Value == "" {
			http.Error(w, "Invalid access token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
