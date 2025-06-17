package internal

import (
	"dubai-auto/internal/config"
	"dubai-auto/internal/route"
	"dubai-auto/pkg"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitApp(db *pgxpool.Pool, conf *config.Config) *gin.Engine {

	if config.ENV.GIN_MODE == "release" {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	}

	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.Use(pkg.Cors)

	router.Static("/images", "."+conf.STATIC_PATH)

	// new routers
	route.Init(router, db)
	return router
}
