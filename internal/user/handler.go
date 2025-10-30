package user

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/Infamous003/go-blog-backend/internal/validate"

	"github.com/Infamous003/go-blog-backend/utils"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	s *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{s: s}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Post("/auth/register", h.handleRegister)
	r.Get("/users/username/{username}", h.handleGetByUsername)
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	payload := &UserRegister{} // creating an empty payload
	if err := utils.FromJSON(r.Body, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := validate.Struct(payload); err != nil {
		errs := strings.Join(validate.ErrorMessages(err), ", ")
		utils.WriteError(w, http.StatusBadRequest, errs)
		return
	}

	user, err := h.s.RegisterUser(r.Context(), payload)
	if err != nil {
		if errors.Is(err, ErrUserAlreadyExists) {
			utils.WriteError(w, http.StatusConflict, "username already exists")
			return
		}
		log.Printf("[ERROR] register user: %v", err)
		utils.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff") // stops browsers from guessing content type
	w.WriteHeader(http.StatusCreated)

	if err := utils.ToJSON(w, user); err != nil {
		log.Printf("[ERROR] ToJSON: %v", err)
		utils.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}
}

func (h *Handler) handleGetByUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	user, err := h.s.GetUser(r.Context(), username)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			utils.WriteError(w, http.StatusNotFound, "user not found")
			return
		}
		log.Printf("[ERROR] get user: %v", err)
		utils.WriteError(w, http.StatusInternalServerError, "internal server error")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)

	if err := utils.ToJSON(w, user); err != nil {
		log.Printf("[ERROR] ToJSON: %v", err)
		utils.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}
}
