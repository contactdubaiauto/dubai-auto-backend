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
	thirdPartyHandler := http.NewThirdPartyHandler(thirdPartyService, auth.NewValidator())

	{
		r.Get("/registration-data", thirdPartyHandler.GetRegistrationData)
		r.Get("/profile", auth.TokenGuard, thirdPartyHandler.GetProfile)
		r.Get("/profile/my-cars", auth.TokenGuard, thirdPartyHandler.GetMyCars)
		r.Get("/profile/on-sale", auth.TokenGuard, thirdPartyHandler.OnSale)
		r.Post("/first-login", auth.TokenGuard, thirdPartyHandler.FirstLogin)
		r.Post("/profile/banner", auth.TokenGuard, thirdPartyHandler.BannerImage)
		r.Post("/profile/images", auth.TokenGuard, thirdPartyHandler.AvatarImages)
		r.Post("/profile", auth.TokenGuard, thirdPartyHandler.Profile)
		// dealer routes
		r.Post("/dealer/car", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.CreateDealerCar)
		r.Post("/dealer/car/:id", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.UpdateDealerCar)
		r.Get("/dealer/car/:id/edit", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.GetEditCarByID)
		r.Post("/dealer/car/:id/images", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.CreateDealerCarImages)
		r.Post("/dealer/car/:id/videos", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.CreateDealerCarVideos)
		r.Post("/dealer/car/:id/sell", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.StatusDealer)
		r.Post("/dealer/car/:id/dont-sell", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.StatusDealer)
		r.Delete("/dealer/car/:id/images", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.DeleteDealerCarImage)
		r.Delete("/dealer/car/:id/videos", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.DeleteDealerCarVideo)
		r.Delete("/dealer/car/:id", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.DeleteDealerCar)
		// logist routes
		r.Get("/logist/destinations", auth.TokenGuard, auth.LogistGuard, thirdPartyHandler.GetLogistDestinations)
		r.Post("/logist/destinations", auth.TokenGuard, auth.LogistGuard, thirdPartyHandler.CreateLogistDestination)
		r.Delete("/logist/destinations/:id", auth.TokenGuard, auth.LogistGuard, thirdPartyHandler.DeleteLogistDestination)
		// broker routes
		// car service routes
	}
}
