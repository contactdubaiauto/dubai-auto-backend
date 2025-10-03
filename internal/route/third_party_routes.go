package route

import (
	"dubai-auto/internal/delivery/http"
	"dubai-auto/internal/repository"
	"dubai-auto/internal/service"
	"dubai-auto/pkg/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupThirdPartyRoutes(r fiber.Router, db *pgxpool.Pool) {
	thirdPartyRepository := repository.NewThirdPartyRepository(db)
	thirdPartyService := service.NewThirdPartyService(thirdPartyRepository)
	thirdPartyHandler := http.NewThirdPartyHandler(thirdPartyService)

	{
		r.Get("/profile", auth.TokenGuard, thirdPartyHandler.GetProfile)
		r.Get("/registration-data", thirdPartyHandler.GetRegistrationData)
		r.Post("/profile", auth.TokenGuard, thirdPartyHandler.Profile)
	}
}
