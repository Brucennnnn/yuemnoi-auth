package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sds-2/feature/auth"
	"github.com/sds-2/feature/item"
	"github.com/sds-2/feature/review"
	"github.com/sds-2/feature/user"
)

type Handler struct {
	itemHandler   *item.ItemHandler
	reviewHandler *review.ReviewHandler
	userHandler   *user.UserHandler
	authHandler   *auth.AuthHandler
}

func NewHandler(itemHandler *item.ItemHandler, reviewHandler *review.ReviewHandler, userHandler *user.UserHandler, authHandler *auth.AuthHandler) *Handler {
	return &Handler{
		itemHandler:   itemHandler,
		reviewHandler: reviewHandler,
		userHandler:   userHandler,
		authHandler:   authHandler,
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
	{
		authRouter := r.Group("/auth")
		authRouter.Get("/google/login", h.authHandler.OAuthLogin)
		authRouter.Get("/google/callback", h.authHandler.OAuthCallback)
	}

	{
		userRouter := r.Group("/user")
		userRouter.Post("/", h.userHandler.CreateUser)
		userRouter.Get("/", h.userHandler.GetUsers)
		userRouter.Get("/:student_id", h.userHandler.ViewUserProfile)
		userRouter.Patch("/:student_id", h.userHandler.ManageUserProfile)
		userRouter.Delete("/:student_id", h.userHandler.DeleteUser)
	}
}
