package handler

import (
	"RazdelyCheck/internal/dto"
	"encoding/json"
	"net/http"

	"RazdelyCheck/internal/service"
	"RazdelyCheck/internal/util"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// POST /users
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name    *string    `json:"name"`
		Email   *string    `json:"email"`
		OwnerID *uuid.UUID `json:"ownerId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	user := &dto.User{
		Name:    body.Name,
		Email:   body.Email,
		OwnerID: body.OwnerID,
	}

	if err := h.service.CreateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	util.RespondJSON(w, http.StatusCreated, user)
}

// PUT /users/{id}
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid userID", http.StatusBadRequest)
		return
	}

	var body struct {
		Name    *string    `json:"name"`
		Email   *string    `json:"email"`
		OwnerID *uuid.UUID `json:"ownerId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	user := &dto.User{
		ID:      id,
		Name:    body.Name,
		Email:   body.Email,
		OwnerID: body.OwnerID,
	}

	if err := h.service.Update(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	util.RespondJSON(w, http.StatusOK, user)
}

// GET /users/owned
func (h *UserHandler) ListOwnedUsers(w http.ResponseWriter, r *http.Request) {
	ownerID := r.Context().Value("userID").(uuid.UUID)

	users, err := h.service.ListOwnedUsers(ownerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, users)
}

// POST /users/{id}/make-permanent
func (h *UserHandler) MakePermanent(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	tempUserID, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid userID", http.StatusBadRequest)
		return
	}

	if err := h.service.MakePermanent(tempUserID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	util.RespondJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// DELETE /users/{id}
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid userID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteUserWithOwned(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}
