package handler

import (
	"net/http"

	"RazdelyCheck/internal/service"
	"RazdelyCheck/internal/util"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type ItemHandler struct {
	service *service.ItemService
}

func NewItemHandler(service *service.ItemService) *ItemHandler {
	return &ItemHandler{service: service}
}

func (h *ItemHandler) ExcludeItem(w http.ResponseWriter, r *http.Request) {
	itemIDStr := chi.URLParam(r, "itemID")
	checkIDStr := r.URL.Query().Get("checkID")

	itemID, err := uuid.Parse(itemIDStr)
	if err != nil {
		util.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid itemID"})
		return
	}
	checkID, err := uuid.Parse(checkIDStr)
	if err != nil {
		util.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid checkID"})
		return
	}

	if err := h.service.ExcludeItem(itemID, checkID); err != nil {
		util.RespondJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	util.RespondJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *ItemHandler) IncludeItem(w http.ResponseWriter, r *http.Request) {
	itemIDStr := chi.URLParam(r, "itemID")
	checkIDStr := r.URL.Query().Get("checkID")

	itemID, err := uuid.Parse(itemIDStr)
	if err != nil {
		util.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid itemID"})
		return
	}
	checkID, err := uuid.Parse(checkIDStr)
	if err != nil {
		util.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid checkID"})
		return
	}

	if err := h.service.IncludeItem(itemID, checkID); err != nil {
		util.RespondJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	util.RespondJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
