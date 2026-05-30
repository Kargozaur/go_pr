package middleware

import (
	"ecommerce/pkg/logger"
	"net/http"
	"time"
)

func ProcessTime(logger *logger.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		message := "request processed in " + duration.String() +
			" for " + r.URL.String()
		logger.Writer.Info("time", "message", message)
	})
}
