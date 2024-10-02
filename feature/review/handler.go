package review

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sds-2/model"
	"strconv"
)

type ReviewHandler struct {
	reviewRepository ReviewRepository
}

func NewReviewHandler(reviewRepository ReviewRepository) *ReviewHandler {
	return &ReviewHandler{
		reviewRepository: reviewRepository, 
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
    reviews, err := h.reviewRepository.GetReviewsByUserId(userId)
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
	return c.Status(200).JSON(response)
}

func (h *ReviewHandler) CreateReview(c *fiber.Ctx) error {
	userIdRaw := c.Locals("user_id") 
	userIdStr, ok := userIdRaw.(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not authenticated",
		})
	}
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID format",
		})
	}
	body := new(CreateReviewRequest)
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}
	if body.Rating < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Rating must be greater than or equal to 0",
		})
	}
	review := model.Review{
		Rating:       body.Rating,
		Description:  body.Description,
		ReviewerID:   userId,
		RevieweeID:   body.RevieweeID,
	}
	err = h.reviewRepository.CreateReview(review)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to create review",
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "create review successfully",
	})
}

