package user

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Post("/register", h.handleRegister)
	r.Post("/login", h.handleLogin)
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Handling Register User"))
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Handling Login User"))
}
