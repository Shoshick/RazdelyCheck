package handler

import (
	"RazdelyCheck/internal/dto"
	"encoding/json"
	"net/http"

	"RazdelyCheck/internal/service"
	"RazdelyCheck/internal/util"

	"github.com/google/uuid"
)

type CheckSourceHandler struct {
	service *service.CheckSourceService
}

func NewCheckSourceHandler(service *service.CheckSourceService) *CheckSourceHandler {
	return &CheckSourceHandler{service: service}
}

// POST /check-source/process-qr
func (h *CheckSourceHandler) ProcessQR(w http.ResponseWriter, r *http.Request) {
	userIDVal := r.Context().Value("userID")
	if userIDVal == nil {
		util.RespondJSON(w, http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
		return
	}
	userID, ok := userIDVal.(uuid.UUID)
	if !ok {
		util.RespondJSON(w, http.StatusInternalServerError, map[string]string{"error": "invalid userID in context"})
		return
	}

	var input struct {
		QRData string `json:"qrData"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		util.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid body"})
		return
	}

	check, err := h.service.ProcessQR(userID, dto.QRScanInput{QRData: input.QRData}, nil)
	if err != nil {
		util.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	util.RespondJSON(w, http.StatusCreated, check)
}
