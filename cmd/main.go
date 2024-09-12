package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/sds-2/cmd/di"
	"github.com/sds-2/config"
	"github.com/sds-2/middleware"
)

func main() {

	ctx := context.Background()
	defer func() {
		if r := recover(); r != nil {
			slog.Error("recover from panic!",
				slog.Any("err", r),
			)
		}
	}()

	cfg := config.Load()
	app := fiber.New()
	handler, err := di.InitDI(ctx, cfg)
	if err != nil {
		slog.Error("failed to initialize DI, exiting...",
			"error", err)
		os.Exit(1)
		return
	}
	app.Use(requestid.New(),
		middleware.SetupUserContext,
		adaptor.HTTPMiddleware(middleware.AuthMiddleware))

	handler.RegisterRouter(app)

	app.Listen(":3000")
}
