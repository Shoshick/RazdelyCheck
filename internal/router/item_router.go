package router

import (
	"RazdelyCheck/internal/handler"
	"github.com/go-chi/chi/v5"
)

func NewItemRouter(r chi.Router, h *handler.ItemHandler) {
	r.Route("/items", func(r chi.Router) {
		r.Post("/{itemID}/exclude", h.ExcludeItem)
		r.Post("/{itemID}/include", h.IncludeItem)
	})
}
