package user

import (
	"errors"
	"log"
	"net/http"

	"github.com/Infamous003/go-blog-backend/internal/validator"
	// "github.com/go-playground/validator/v10"

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
	r.Get("/users/{id}", h.handleGetByID)
	r.Delete("/users/{id}", h.handleDelete)
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	payload := &UserRegister{} // creating an empty payload
	if err := utils.ReadJSON(w, r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	v := validator.New()
	ValidateUser(v, payload)
	if !v.Valid() {
		utils.FailedValidationResponse(w, r, v.Errors)
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

	headers := make(http.Header)
	headers.Set("X-Content-Type-Options", "nosniff")

	if err := utils.WriteJSON(w, http.StatusCreated, user, headers); err != nil {
		log.Printf("[ERROR] WriteJSON: %v", err)
	}
}

func (h *Handler) handleGetByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromURL(r)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, "not found")
		return
	}

	user, err := h.s.GetUser(r.Context(), id)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			utils.WriteError(w, http.StatusNotFound, "user not found")
			return
		}
		log.Printf("[ERROR] get user: %v", err)
		utils.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	if err = utils.WriteJSON(w, http.StatusOK, user, nil); err != nil {
		log.Printf("[ERROR] ToJSON: %v", err)
		return
	}
}

func (h *Handler) handleDelete(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromURL(r)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, "not found")
		return
	}

	err = h.s.DeleteByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			utils.WriteError(w, http.StatusNotFound, "user not found")
			return
		}
		utils.WriteError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
}
