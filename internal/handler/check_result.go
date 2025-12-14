package handler

import (
	"encoding/json"
	"net/http"

	"RazdelyCheck/internal/service"
	"RazdelyCheck/internal/util"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type CheckResultHandler struct {
	service *service.CheckResultService
}

func NewCheckResultHandler(service *service.CheckResultService) *CheckResultHandler {
	return &CheckResultHandler{service: service}
}

// POST /mini-check/{checkID}/owner/{ownerID}
func (h *CheckResultHandler) CreateOwnerMiniCheck(w http.ResponseWriter, r *http.Request) {
	checkID, err := uuid.Parse(chi.URLParam(r, "checkID"))
	if err != nil {
		http.Error(w, "invalid checkID", http.StatusBadRequest)
		return
	}
	ownerID, err := uuid.Parse(chi.URLParam(r, "ownerID"))
	if err != nil {
		http.Error(w, "invalid ownerID", http.StatusBadRequest)
		return
	}

	cr, err := h.service.CreateOwnerMiniCheck(checkID, ownerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusCreated, cr)
}

// POST /mini-check/add
func (h *CheckResultHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	var body struct {
		CheckResultID string  `json:"checkResultID"`
		ItemID        string  `json:"itemID"`
		Quantity      float64 `json:"quantity"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	checkResultID, _ := uuid.Parse(body.CheckResultID)
	itemID, _ := uuid.Parse(body.ItemID)

	if err := h.service.AddItemToCheckResult(checkResultID, itemID, body.Quantity); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// POST /mini-check/remove
func (h *CheckResultHandler) RemoveItem(w http.ResponseWriter, r *http.Request) {
	var body struct {
		CheckResultID string `json:"checkResultID"`
		ItemID        string `json:"itemID"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	checkResultID, _ := uuid.Parse(body.CheckResultID)
	itemID, _ := uuid.Parse(body.ItemID)

	if err := h.service.RemoveItemFromCheckResult(checkResultID, itemID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// POST /mini-check/update
func (h *CheckResultHandler) UpdateQuantity(w http.ResponseWriter, r *http.Request) {
	var body struct {
		CheckResultID string  `json:"checkResultID"`
		ItemID        string  `json:"itemID"`
		Quantity      float64 `json:"quantity"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	checkResultID, _ := uuid.Parse(body.CheckResultID)
	itemID, _ := uuid.Parse(body.ItemID)

	if err := h.service.UpdateItemQuantityInCheckResult(checkResultID, itemID, body.Quantity); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// GET /mini-check/{checkResultID}/items
func (h *CheckResultHandler) GetCheckItems(w http.ResponseWriter, r *http.Request) {
	checkResultID, err := uuid.Parse(chi.URLParam(r, "checkResultID"))
	if err != nil {
		http.Error(w, "invalid checkResultID", http.StatusBadRequest)
		return
	}

	items, err := h.service.GetCheckItems(checkResultID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, items)
}

// GET /mini-check/{checkID}/all
func (h *CheckResultHandler) GetAllMiniChecks(w http.ResponseWriter, r *http.Request) {
	checkID, err := uuid.Parse(chi.URLParam(r, "checkID"))
	if err != nil {
		http.Error(w, "invalid checkID", http.StatusBadRequest)
		return
	}

	results, err := h.service.GetAllMiniChecks(checkID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, http.StatusOK, results)
}
