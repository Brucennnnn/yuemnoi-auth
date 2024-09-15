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

func (h *ReviewHandler) ViewUserReviews(c *fiber.Ctx) error {
	userIdStr := c.Params("userId")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}
    reviews, err := h.reviewDomain.ViewUserReviews(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to retrieve reviews",
		})
	}
	var response []ViewUserReviewsResponse
	for _, review := range reviews {
		response = append(response, ViewUserReviewsResponse{
			ID: review.ID,
			Rating:review.Rating,
			Description:review.Description,
			ReviewerID:review.ReviewerID,
		})
	}
	return c.JSON(response)
}

func (h *ReviewHandler) ReviewUser(c *fiber.Ctx) error {
	body := new(ReviewUserRequest)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}
	review := model.Review{
		Rating:       body.Rating,
		Description:  body.Description,
		ReviewerID:   body.ReviewerID,
		RevieweeID:   body.RevieweeID,
	}
	res, err := h.reviewDomain.ReviewUser(review)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to create review",
		})
	}
	return c.JSON(res)
}