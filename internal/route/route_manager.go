package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"

	_ "dubai-auto/docs"
	"dubai-auto/pkg/auth"
)

func Init(app *fiber.App, db *pgxpool.Pool) {

	userRoute := app.Group("/api/v1/users")
	SetupUserRoutes(userRoute, db)

	authRoute := app.Group("/api/v1/auth")
	SetupAuthRoutes(authRoute, db)

	motorcycleRoute := app.Group("/api/v1/motorcycles")
	SetupMotorcycleRoutes(motorcycleRoute, db)

	comtransRoute := app.Group("/api/v1/comtrans")
	SetupComtranRoutes(comtransRoute, db)

	adminRoute := app.Group("/api/v1/admin", auth.TokenGuard, auth.AdminGuard)
	SetupAdminRoutes(adminRoute, db)

	thirdPartyRoute := app.Group("/api/v1/third-party")
	SetupThirdPartyRoutes(thirdPartyRoute, db)

	SetupWebSocketRoutes(app, db)

}
