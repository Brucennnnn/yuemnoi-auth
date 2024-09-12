package items

import (
	"github.com/gofiber/fiber/v2"
)

type ItemsHandler struct {
	ItemsDomain ItemsDomain
}

func NewItemsHandler(itemsDomain ItemsDomain) *ItemsHandler {
	return &ItemsHandler{
		ItemsDomain: itemsDomain,
	}
}

func (h *ItemsHandler) GetAllItems(c *fiber.Ctx) error {
	res, err := h.ItemsDomain.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to retrieve items",
		})
	}
	return c.JSON(res)
}
