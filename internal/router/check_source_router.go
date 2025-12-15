package router

import (
	"RazdelyCheck/internal/handler"

	"github.com/go-chi/chi/v5"
)

func NewCheckSourceRouter(r chi.Router, h *handler.CheckSourceHandler) {
	r.Route("/check-sources", func(r chi.Router) {
		r.Post("/process-qr", h.ProcessQR)
	})
}
