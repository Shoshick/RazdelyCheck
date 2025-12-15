package router

import (
	"RazdelyCheck/internal/handler"

	"github.com/go-chi/chi/v5"
)

func NewGroupRouter(r chi.Router, h *handler.GroupHandler) {
	r.Route("/groups", func(r chi.Router) {
		// CRUD
		r.Post("/", h.CreateGroup)
		r.Get("/", h.ListGroups)
		r.Get("/{id}", h.GetGroup)
		r.Put("/{id}", h.UpdateGroup)
		r.Delete("/{id}", h.DeleteGroup)

		// Users in group
		r.Post("/{id}/users", h.AddUser)
		r.Delete("/{id}/users/{userID}", h.RemoveUser)
		r.Get("/{id}/users", h.ListUsersByGroup)

		// Groups by user
		r.Get("/user/{userID}", h.ListGroupsByUser)
	})
}
