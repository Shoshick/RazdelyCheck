package router

import (
	"RazdelyCheck/internal/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewItemRouter(h *handler.ItemHandler) http.Handler {
	r := chi.NewRouter()

	r.Post("/{itemID}/exclude", h.ExcludeItem)
	r.Post("/{itemID}/include", h.IncludeItem)

	return r
}
