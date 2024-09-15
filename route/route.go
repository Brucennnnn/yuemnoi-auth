package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sds-2/feature/item"
)

type Handler struct {
	itemHandler *item.ItemHandler
}

func NewHandler(itemHandler *item.ItemHandler) *Handler {
	return &Handler{
		itemHandler: itemHandler,
	}
}
func (h *Handler) RegisterRouter(r fiber.Router) {
	{
		itemRouter := r.Group("/item")
		itemRouter.Get("/", h.itemHandler.GetAllitem)
	}
}
