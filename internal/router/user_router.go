package router

import (
	"RazdelyCheck/internal/handler"

	"github.com/go-chi/chi/v5"
)

func NewUserRouter(h *handler.UserHandler) chi.Router {
	r := chi.NewRouter()

	r.Post("/", h.CreateUser)                       // POST /users
	r.Put("/{id}", h.UpdateUser)                    // PUT /users/{id}
	r.Get("/owned", h.ListOwnedUsers)               // GET /users/owned
	r.Post("/{id}/make-permanent", h.MakePermanent) // POST /users/{id}/make-permanent
	r.Delete("/{id}", h.DeleteUser)                 // DELETE /users/{id}

	return r
}
