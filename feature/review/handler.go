package review

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sds-2/model"
	"strconv"
)

type ReviewHandler struct {
	reviewDomain ReviewDomain
}

func NewReviewHandler(reviewDomain ReviewDomain) *ReviewHandler {
	return &ReviewHandler{
		reviewDomain: reviewDomain,
	}
}

func (h *ReviewHandler) GetReviewsByUserId(c *fiber.Ctx) error {
	userIdStr := c.Params("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}
    reviews, err := h.reviewDomain.GetReviewsByUserId(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to retrieve reviews",
		})
	}
	var response []GetReviewsByUserIdResponse
	for _, review := range reviews {
		response = append(response, GetReviewsByUserIdResponse{
			ID: review.ID,
			Rating:review.Rating,
			Description:review.Description,
			ReviewerID:review.ReviewerID,
		})
	}
	return c.JSON(response)
}

func (h *ReviewHandler) CreateReview(c *fiber.Ctx) error {
	body := new(CreateReviewRequest)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}
	if body.Rating <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Rating must be positive",
		})
	}
	review := model.Review{
		Rating:       body.Rating,
		Description:  body.Description,
		ReviewerID:   body.ReviewerID,
		RevieweeID:   body.RevieweeID,
	}
	res, err := h.reviewDomain.CreateReview(review)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to create review",
		})
	}
	return c.JSON(res)
}