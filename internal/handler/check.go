package handler

import (
	"encoding/json"
	"net/http"

	"RazdelyCheck/internal/service"
	"RazdelyCheck/internal/util"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type CheckHandler struct {
	service *service.CheckService
}

func NewCheckHandler(service *service.CheckService) *CheckHandler {
	return &CheckHandler{service: service}
}

// POST /checks
func (h *CheckHandler) CreateCheck(w http.ResponseWriter, r *http.Request) {
	var body struct {
		GroupID  *uuid.UUID `json:"groupId"`
		TotalSum int64      `json:"totalSum"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("userID").(uuid.UUID)
	check, err := h.service.Create(userID, body.GroupID, body.TotalSum)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusCreated, check)
}

// DELETE /checks/{id}
func (h *CheckHandler) DeleteCheck(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	checkID, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid checkID", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("userID").(uuid.UUID)
	isAdmin := r.Context().Value("isAdmin").(bool)

	if err := h.service.Delete(checkID, userID, isAdmin); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})
	if err != nil {
		return
	}
}

func (h *CheckHandler) ListChecks(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(uuid.UUID)

	checks, err := h.service.ListByUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, checks)
}
