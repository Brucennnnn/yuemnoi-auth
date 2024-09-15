package item

import (
	"github.com/gofiber/fiber/v2"
)

type ItemHandler struct {
	itemDomain ItemDomain
}

func NewItemHandler(itemDomain ItemDomain) *ItemHandler {
	return &ItemHandler{
		itemDomain: itemDomain,
	}
}

func (h *ItemHandler) GetAllitem(c *fiber.Ctx) error {
	res, err := h.itemDomain.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to retrieve item",
		})
	}
	return c.JSON(res)
}
