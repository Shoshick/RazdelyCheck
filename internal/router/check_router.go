package router

import (
	"RazdelyCheck/internal/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewCheckRouter(h *handler.CheckHandler) http.Handler {
	r := chi.NewRouter()
	r.Post("/", h.CreateCheck)
	r.Delete("/{id}", h.DeleteCheck)
	r.Get("/", h.ListChecks)
	return r
}
