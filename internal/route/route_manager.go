package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"

	_ "dubai-auto/docs"
	"dubai-auto/internal/config"
	"dubai-auto/pkg/auth"
	"dubai-auto/pkg/firebase"
)

func Init(app *fiber.App, config *config.Config, db *pgxpool.Pool, firebaseService *firebase.FirebaseService, validator *auth.Validator) {
	api := app.Group("/api")
	{
		v1 := api.Group("/v1")
		{

			userRoute := v1.Group("/users")
			SetupUserRoutes(userRoute, config, db, validator)

			authRoute := v1.Group("/auth")
			SetupAuthRoutes(authRoute, config, db, validator)

			// motorcycleRoute := v1.Group("/motorcycles")
			// SetupMotorcycleRoutes(motorcycleRoute, config, db, validator)

			// comtransRoute := v1.Group("/comtrans")
			// SetupComtranRoutes(comtransRoute, config, db, validator)

			adminRoute := v1.Group("/admin", auth.TokenGuard, auth.AdminGuard)
			SetupAdminRoutes(adminRoute, config, db, firebaseService, validator)

			thirdPartyRoute := v1.Group("/third-party")
			SetupThirdPartyRoutes(thirdPartyRoute, config, db, validator)

		}
	}

	SetupWebSocketRoutes(app, db, firebaseService, config)

}
