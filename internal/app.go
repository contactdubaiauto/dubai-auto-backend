package internal

import (
	"os"

	"github.com/gofiber/fiber/v2"
	fb_logger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/jackc/pgx/v5/pgxpool"

	"dubai-auto/internal/config"
	"dubai-auto/internal/delivery/http/middleware"
	"dubai-auto/internal/route"
	"dubai-auto/pkg/auth"
	"dubai-auto/pkg/logger"
)

func InitApp(db *pgxpool.Pool, conf *config.Config, logger *logger.Logger) *fiber.App {
	appConfig := fiber.Config{
		BodyLimit: 50 * 1024 * 1024,
	}
	app := fiber.New(appConfig)
	app.Use(pprof.New())
	app.Use(auth.Cors)

	if config.ENV.APP_MODE == "release" {
		app.Use(middleware.ZerologMiddleware(logger))
	} else {
		app.Use(fb_logger.New(fb_logger.Config{
			Format: "[${time}] ${ip} ${status} - ${method} ${path} ${latency}\n",
			Output: os.Stdout, // Or to a file: os.OpenFile("fiber.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		}))
	}

	app.Static("/images", "."+conf.STATIC_PATH)

	route.Init(app, db)
	return app
}
