package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"

	_ "dubai-auto/docs"
	"dubai-auto/pkg/auth"
)

func Init(r *fiber.App, db *pgxpool.Pool) {

	userRoute := r.Group("/api/v1/users")
	SetupUserRoutes(userRoute, db)

	authRoute := r.Group("/api/v1/auth")
	SetupAuthRoutes(authRoute, db)

	motorcycleRoute := r.Group("/api/v1/motorcycles")
	SetupMotorcycleRoutes(motorcycleRoute, db)

	comtransRoute := r.Group("/api/v1/comtrans")
	SetupComtranRoutes(comtransRoute, db)

	adminRoute := r.Group("/api/v1/admin", auth.TokenGuard)
	SetupAdminRoutes(adminRoute, db)

	thirdPartyRoute := r.Group("/api/v1/third-party")
	SetupThirdPartyRoutes(thirdPartyRoute, db)

}
