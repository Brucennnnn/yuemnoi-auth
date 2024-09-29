package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sds-2/config"
	"github.com/sds-2/feature/auth"
	"github.com/sds-2/feature/item"
	"github.com/sds-2/feature/review"
	"github.com/sds-2/feature/user"
	"github.com/sds-2/middleware"
)

type Handler struct {
	itemHandler   *item.ItemHandler
	reviewHandler *review.ReviewHandler
	authHandler   *auth.AuthHandler
	userHandler   *user.UserHandler
}

func NewHandler(itemHandler *item.ItemHandler, reviewHandler *review.ReviewHandler, authHandler *auth.AuthHandler, userHandler *user.UserHandler) *Handler {
	return &Handler{
		itemHandler:   itemHandler,
		reviewHandler: reviewHandler,
		authHandler:   authHandler,
		userHandler:   userHandler,
	}
}

func (h *Handler) RegisterRouter(r fiber.Router, cfg *config.Config) {
	{
		itemRouter := r.Group("/item", middleware.AuthMiddleware(cfg))
		itemRouter.Get("/", h.itemHandler.GetAllitem)
	}
	{
		reviewRouter := r.Group("/review", middleware.AuthMiddleware(cfg))
		reviewRouter.Get("/user/:userId", h.reviewHandler.GetReviewsByUserId)
		reviewRouter.Post("/", h.reviewHandler.CreateReview)
	}
	{
		authRouter := r.Group("/auth")
		authRouter.Get("/google/login", h.authHandler.OAuthLogin)
		authRouter.Get("/google/callback", h.authHandler.OAuthCallback)
	}
	{
		userRouter := r.Group("/user", middleware.AuthMiddleware(cfg))
		userRouter.Get("/", h.userHandler.GetUsers)
		userRouter.Get("/:id", h.userHandler.GetUserById)
		userRouter.Patch("/:id", h.userHandler.UpdateUser)
		userRouter.Delete("/:id", h.userHandler.DeleteUser)
		userRouter.Get("/email/:email", h.userHandler.GetUserByEmail)
	}
}
