package router

import (
	"RazdelyCheck/internal/handler"

	"github.com/go-chi/chi/v5"
)

func NewUserRouter(r chi.Router, h *handler.UserHandler) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", h.CreateUser)
		r.Put("/{id}", h.UpdateUser)
		r.Get("/owned", h.ListOwnedUsers)
		r.Post("/{id}/make-permanent", h.MakePermanent)
		r.Delete("/{id}", h.DeleteUser)
	})
}
