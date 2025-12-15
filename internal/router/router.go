package router

import (
	"net/http"

	"RazdelyCheck/internal/handler"
	"RazdelyCheck/internal/middleware"

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

	r.Use(middleware.Logging)
	r.Use(middleware.Recovery)

	NewCheckRouter(r, checkHandler)
	NewCheckResultRouter(r, checkResultHandler)
	NewItemRouter(r, itemHandler)
	NewGroupRouter(r, groupHandler)
	NewUserRouter(r, userHandler)
	NewCheckSourceRouter(r, checkSourceHandler)

	return r
}
