package route

import (
	"github.com/gofiber/fiber/v2"

	"github.com/sds-2/feature/items"
)

type Handler struct {
	ItemsHandler *items.ItemsHandler
}

func NewHandler(ItemsHandler *items.ItemsHandler) *Handler {
	return &Handler{
		ItemsHandler: ItemsHandler,
	}
}
func (h *Handler) RegisterRouter(r fiber.Router) {
	{
		itemsRouter := r.Group("/items")
		itemsRouter.Get("/", h.ItemsHandler.GetAllItems)
	}
}
