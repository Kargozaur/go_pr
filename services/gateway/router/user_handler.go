package router

import (
	"context"
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
	ctx, cancel := context.WithTimeout(r.Context(),
		time.Second*5)
	defer cancel()
	link := fmt.Sprintf("http://%s:8082/users/register",
		os.Getenv("HOST"))
	req, err := http.NewRequestWithContext(ctx,
		http.MethodPost,
		link,
		r.Body)
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
	body, err := json.Marshal(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

func (h *Handler) loginUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(),
		time.Second*5)
	defer cancel()
	link := fmt.Sprintf("http://%s:8082/users/login",
		os.Getenv("HOST"))
	req, err := http.NewRequestWithContext(ctx,
		http.MethodPost,
		link,
		r.Body)
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
	w.Header().Set("Content-Type", "application/json")
	body, err := json.Marshal(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
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
	w.Header().Set("Content-Type", "application/json")
	body, err := json.Marshal(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(body)
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
	w.Header().Set("Content-Type", "application/json")
	body, err := json.Marshal(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(body)
}

func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	accessCookie, refreshCookie, err := prepareCookie(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	link := fmt.Sprintf("http://%s:8082/users/profile/me",
		os.Getenv("HOST"))
	ctx, cancel := context.WithTimeout(r.Context(),
		time.Second*5)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx,
		http.MethodGet,
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
	w.Header().Set("Content-Type", "application/json")
	body, err := json.Marshal(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(body)
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
