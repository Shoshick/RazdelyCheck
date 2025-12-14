package router

import (
	"net/http"

	"RazdelyCheck/internal/handler"

	"github.com/go-chi/chi/v5"
)

// NewRouter собирает все подроутеры проекта
func NewRouter(
	checkHandler *handler.CheckHandler,
	itemHandler *handler.ItemHandler,
	groupHandler *handler.GroupHandler,
	userHandler *handler.UserHandler,
	checkSourceHandler *handler.CheckSourceHandler,
	checkResultHandler *handler.CheckResultHandler,
) http.Handler {
	r := chi.NewRouter()

	// middleware можно добавить сюда, если нужно
	// r.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)

	// подроутеры
	r.Route("/checks", func(r chi.Router) {
		r.Mount("/", NewCheckRouter(checkHandler))
	})

	r.Route("/check-results", func(r chi.Router) {
		r.Mount("/", NewCheckResultRouter(checkResultHandler))
	})

	r.Route("/items", func(r chi.Router) {
		r.Mount("/", NewItemRouter(itemHandler))
	})

	r.Route("/groups", func(r chi.Router) {
		r.Mount("/", NewGroupRouter(groupHandler))
	})

	r.Route("/users", func(r chi.Router) {
		r.Mount("/", NewUserRouter(userHandler))
	})

	r.Route("/check_sources", func(r chi.Router) {
		r.Mount("/", NewCheckSourceRouter(checkSourceHandler))
	})

	return r
}
