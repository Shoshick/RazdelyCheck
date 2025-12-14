package router

import (
	"RazdelyCheck/internal/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewCheckResultRouter(h *handler.CheckResultHandler) http.Handler {
	r := chi.NewRouter()

	// Создать миничек владельца
	r.Post("/{checkID}/owner/{ownerID}", h.CreateOwnerMiniCheck)

	// Добавить предмет в миничек
	r.Post("/add", h.AddItem)

	// Удалить предмет из миничека
	r.Post("/remove", h.RemoveItem)

	// Обновить количество предмета в миничеке
	r.Post("/update", h.UpdateQuantity)

	// Получить все предметы миничека
	r.Get("/{checkResultID}/items", h.GetCheckItems)

	// Получить список всех миничеков по чеку
	r.Get("/{checkID}/all", h.GetAllMiniChecks)

	return r
}
