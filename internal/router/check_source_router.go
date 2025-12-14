package router

import (
	"RazdelyCheck/internal/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewCheckSourceRouter(h *handler.CheckSourceHandler) http.Handler {
	r := chi.NewRouter()
	r.Post("/process-qr", h.ProcessQR)
	return r
}
