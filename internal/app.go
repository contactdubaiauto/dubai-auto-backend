package internal

import (
	"dubai-auto/internal/config"
	"dubai-auto/internal/delivery/http/middleware"
	"dubai-auto/internal/route"
	"dubai-auto/pkg"
	"dubai-auto/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitApp(db *pgxpool.Pool, conf *config.Config, logger *logger.Logger) *fiber.App {
	appConfig := fiber.Config{
		BodyLimit: 50 * 1024 * 1024,
	}
	app := fiber.New(appConfig)
	app.Use(pkg.Cors)

	if config.ENV.GIN_MODE == "release" {
		app.Use(middleware.ZerologMiddleware(logger))
	}

	app.Static("/images", "."+conf.STATIC_PATH)

	route.Init(app, db)
	return app
}
