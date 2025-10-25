package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"

	_ "dubai-auto/docs"
	"dubai-auto/internal/config"
	"dubai-auto/pkg/auth"
	"dubai-auto/pkg/firebase"
)

func Init(app *fiber.App, config *config.Config, db *pgxpool.Pool, firebaseService *firebase.FirebaseService) {

	userRoute := app.Group("/api/v1/users")
	SetupUserRoutes(userRoute, config, db)

	authRoute := app.Group("/api/v1/auth")
	SetupAuthRoutes(authRoute, config, db)

	motorcycleRoute := app.Group("/api/v1/motorcycles")
	SetupMotorcycleRoutes(motorcycleRoute, config, db)

	comtransRoute := app.Group("/api/v1/comtrans")
	SetupComtranRoutes(comtransRoute, config, db)

	adminRoute := app.Group("/api/v1/admin", auth.TokenGuard, auth.AdminGuard)
	SetupAdminRoutes(adminRoute, config, db)

	thirdPartyRoute := app.Group("/api/v1/third-party")
	SetupThirdPartyRoutes(thirdPartyRoute, config, db)

	SetupWebSocketRoutes(app, db, firebaseService)

}
