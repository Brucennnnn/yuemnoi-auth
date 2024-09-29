package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
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
		middleware.SetupUserContext)

	handler.RegisterRouter(app, cfg)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0" // Default host if not specified
	}

	app.Listen(host + ":" + port)
}
