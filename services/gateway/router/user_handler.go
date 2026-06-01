package router

import (
	"bytes"
	"context"
	"ecommerce/gateway/schemas"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Handler struct {
	client *http.Client
}

func NewHandler() *Handler {
	return &Handler{
		client: &http.Client{},
	}
}

func (h *Handler) registerUser(w http.ResponseWriter, r *http.Request) {
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
	ctx, cancel := context.WithTimeout(r.Context(),
		time.Second*5)
	defer cancel()
	link := fmt.Sprintf("http://%s:8082/users/register",
		os.Getenv("HOST"))
	req, err := http.NewRequestWithContext(ctx,
		http.MethodPost,
		link,
		bytes.NewBuffer(userBytes))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := h.client.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			http.Error(w, "request timed out", http.StatusRequestTimeout)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		http.Error(w, resp.Status, resp.StatusCode)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(userBytes)
}

func (h *Handler) loginUser(w http.ResponseWriter, r *http.Request) {
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
	ctx, cancel := context.WithTimeout(r.Context(),
		time.Second*5)
	defer cancel()
	link := fmt.Sprintf("http://%s:8082/users/login",
		os.Getenv("HOST"))
	req, err := http.NewRequestWithContext(ctx,
		http.MethodPost,
		link,
		bytes.NewBuffer(loginBytes))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp, err := h.client.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			http.Error(w, "request timed out", http.StatusRequestTimeout)
			return
		}
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bodyBytes)
}

func (h *Handler) logoutUserSingle(w http.ResponseWriter, r *http.Request) {
	accessCookie, refreshCookie, err := prepareCookie(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	link := fmt.Sprintf("http://%s:8082/users/logout/single", os.Getenv("HOST"))
	ctx, cancel := context.WithTimeout(r.Context(),
		time.Second*5)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx,
		http.MethodPost,
		link,
		nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.AddCookie(accessCookie)
	req.AddCookie(refreshCookie)
	req.Header.Set("Authorization",
		fmt.Sprintf("Bearer %s", accessCookie.Value))
	resp, err := h.client.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			http.Error(w, "request timed out", http.StatusRequestTimeout)
			return
		}
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bodyBytes)
}

func (h *Handler) logoutUserAll(w http.ResponseWriter, r *http.Request) {
	accessCookie, refreshCookie, err := prepareCookie(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	link := fmt.Sprintf("http://%s:8082/users/logout/all",
		os.Getenv("HOST"))
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Second*5)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx,
		http.MethodPost,
		link,
		nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.AddCookie(accessCookie)
	req.AddCookie(refreshCookie)
	req.Header.Set("Authorization",
		fmt.Sprintf("Bearer %s", accessCookie.Value))
	resp, err := h.client.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			http.Error(w, "request timed out", http.StatusRequestTimeout)
			return
		}
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bodyBytes)
}

func prepareCookie(r *http.Request) (*http.Cookie, *http.Cookie, error) {
	accessCookie, err := r.Cookie("access_token")
	if err != nil {
		return nil, nil, err
	}
	refreshCookie, err := r.Cookie("refresh_token")
	if err != nil {
		return nil, nil, err
	}
	return accessCookie, refreshCookie, nil
}
