package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sds-2/feature/item"
	"github.com/sds-2/feature/review"
)

type Handler struct {
	itemHandler *item.ItemHandler
	reviewHandler *review.ReviewHandler
}

func NewHandler(itemHandler *item.ItemHandler , reviewHandler *review.ReviewHandler) *Handler {
	return &Handler{
		itemHandler: itemHandler,
		reviewHandler: reviewHandler,
	}
}
func (h *Handler) RegisterRouter(r fiber.Router) {
	{
		itemRouter := r.Group("/item")
		itemRouter.Get("/", h.itemHandler.GetAllitem)
	}
	{
		reviewRouter := r.Group("/review")
		reviewRouter.Get("/user/:userId", h.reviewHandler.GetReviewsByUserId) 
		reviewRouter.Post("/", h.reviewHandler.CreateReview)             
	}
}