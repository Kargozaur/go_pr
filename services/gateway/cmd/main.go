package main

import (
	"ecommerce/gateway/middleware"
	"ecommerce/pkg/logger"
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	logger, err := logger.NewLogWriter("gateway")
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Close()
	mux := http.NewServeMux()
	timeWrapped := middleware.ProcessTime(logger, mux)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{"message": "default message"}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.ListenAndServe(":8080", timeWrapped)
}
