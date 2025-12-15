package router

import (
	"RazdelyCheck/internal/handler"
	"github.com/go-chi/chi/v5"
)

func NewCheckResultRouter(r chi.Router, h *handler.CheckResultHandler) {
	r.Route("/check-results", func(r chi.Router) {
		r.Post("/{checkID}/owner/{ownerID}", h.CreateOwnerMiniCheck)
		r.Post("/add", h.AddItem)
		r.Post("/remove", h.RemoveItem)
		r.Post("/update", h.UpdateQuantity)
		r.Get("/{checkResultID}/items", h.GetCheckItems)
		r.Get("/{checkID}/all", h.GetAllMiniChecks)
	})
}
