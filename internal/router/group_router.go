package router

import (
	"RazdelyCheck/internal/handler"

	"github.com/go-chi/chi/v5"
	"net/http"
)

func NewGroupRouter(h *handler.GroupHandler) http.Handler {
	r := chi.NewRouter()

	// CRUD группы
	r.Post("/", h.CreateGroup)       // POST /groups
	r.Get("/", h.ListGroups)         // GET /groups
	r.Get("/{id}", h.GetGroup)       // GET /groups/{id}
	r.Put("/{id}", h.UpdateGroup)    // PUT /groups/{id}
	r.Delete("/{id}", h.DeleteGroup) // DELETE /groups/{id}

	// Пользователи в группе
	r.Post("/{id}/users", h.AddUser)               // POST /groups/{id}/users
	r.Delete("/{id}/users/{userID}", h.RemoveUser) // DELETE /groups/{id}/users/{userID}
	r.Get("/{id}/users", h.ListUsersByGroup)       // GET /groups/{id}/users

	// Списки групп по пользователю
	r.Get("/user/{userID}", h.ListGroupsByUser) // GET /groups/user/{userID}

	return r
}
