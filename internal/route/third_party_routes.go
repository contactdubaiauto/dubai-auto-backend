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
		r.Get("/registration-data", thirdPartyHandler.GetRegistrationData)
		r.Post("/first-login", auth.TokenGuard, thirdPartyHandler.FirstLogin)
		r.Get("/profile", auth.TokenGuard, thirdPartyHandler.GetProfile)
		r.Post("/profile", auth.TokenGuard, thirdPartyHandler.Profile)
		r.Post("/profile/banner", auth.TokenGuard, thirdPartyHandler.BannerImage)
		r.Post("/profile/images", auth.TokenGuard, thirdPartyHandler.AvatarImages)
	}
}
