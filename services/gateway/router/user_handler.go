package router

import (
	"bytes"
	"context"
	"ecommerce/gateway/schemas"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func registerUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	var userStruct schemas.User
	if err := json.NewDecoder(r.Body).Decode(&userStruct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userBytes, err := json.Marshal(userStruct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*5))
	link := fmt.Sprintf("http://%s:8002/users/register", os.Getenv("HOST"))
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "POST", link, bytes.NewBuffer(userBytes))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		http.Error(w, resp.Status, resp.StatusCode)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(userBytes)
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	var loginStruct schemas.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginStruct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	loginBytes, err := json.Marshal(loginStruct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*5))
	link := fmt.Sprintf("http://%s:8002/users/login", os.Getenv("HOST"))
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "POST", link, bytes.NewBuffer(loginBytes))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		http.Error(w, resp.Status, resp.StatusCode)
		return
	}
	bodyBytes, err := json.Marshal(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bodyBytes)
}
