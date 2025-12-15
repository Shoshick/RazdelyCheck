package router

import (
	"RazdelyCheck/internal/handler"

	"github.com/go-chi/chi/v5"
)

func NewCheckRouter(r chi.Router, h *handler.CheckHandler) {
	r.Route("/checks", func(r chi.Router) {
		r.Post("/", h.CreateCheck)
		r.Get("/", h.ListChecks)
		r.Delete("/{id}", h.DeleteCheck)
	})
}
