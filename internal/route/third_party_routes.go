package route

import (
	"dubai-auto/internal/config"
	"dubai-auto/internal/delivery/http"
	"dubai-auto/internal/repository"
	"dubai-auto/internal/service"
	"dubai-auto/pkg/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupThirdPartyRoutes(r fiber.Router, config *config.Config, db *pgxpool.Pool, validator *auth.Validator) {
	thirdPartyRepository := repository.NewThirdPartyRepository(config, db)
	thirdPartyService := service.NewThirdPartyService(thirdPartyRepository)
	thirdPartyHandler := http.NewThirdPartyHandler(thirdPartyService, validator)

	{
		r.Get("/registration-data", auth.LanguageChecker, thirdPartyHandler.GetRegistrationData)
		r.Get("/profile", auth.TokenGuard, auth.ThirdPartyGuard, auth.LanguageChecker, thirdPartyHandler.GetProfile)
		r.Get("/profile/my-cars", auth.TokenGuard, auth.ThirdPartyGuard, auth.LanguageChecker, thirdPartyHandler.GetMyCars)
		r.Get("/profile/on-sale", auth.TokenGuard, auth.ThirdPartyGuard, auth.LanguageChecker, thirdPartyHandler.OnSale)
		r.Post("/first-login", auth.TokenGuard, auth.ThirdPartyGuard, thirdPartyHandler.FirstLogin)
		r.Post("/profile/banner", auth.TokenGuard, auth.ThirdPartyGuard, thirdPartyHandler.BannerImage)
		r.Delete("/profile/banner", auth.TokenGuard, auth.ThirdPartyGuard, thirdPartyHandler.DeleteBannerImage)
		r.Post("/profile/images", auth.TokenGuard, auth.ThirdPartyGuard, thirdPartyHandler.AvatarImages)
		r.Delete("/profile/images", auth.TokenGuard, auth.ThirdPartyGuard, thirdPartyHandler.DeleteAvatarImages)
		r.Post("/profile", auth.TokenGuard, auth.ThirdPartyGuard, thirdPartyHandler.Profile)

		// dealer car routes
		r.Post("/dealer/car", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.CreateDealerCar)
		r.Get("/dealer/car/:id/edit", auth.TokenGuard, auth.DealerGuard, auth.LanguageChecker, thirdPartyHandler.GetEditCarByID)
		r.Post("/dealer/car/:id/sell", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.StatusDealer)
		r.Post("/dealer/car/:id/dont-sell", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.StatusDealer)
		r.Post("/dealer/car/:id", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.UpdateDealerCar)
		r.Delete("/dealer/car/:id", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.DeleteDealerCar)
		r.Post("/dealer/car/:id/images", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.CreateDealerCarImages)
		r.Post("/dealer/car/:id/videos", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.CreateDealerCarVideos)
		r.Delete("/dealer/car/:id/images", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.DeleteDealerCarImage)
		r.Delete("/dealer/car/:id/videos", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.DeleteDealerCarVideo)

		// dealer motorcycle routes
		r.Post("/dealer/motorcycle", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.CreateDealerMotorcycle)
		r.Get("/dealer/motorcycle/:id/edit", auth.TokenGuard, auth.DealerGuard, auth.LanguageChecker, thirdPartyHandler.GetEditDealerMotorcycleByID)
		r.Post("/dealer/motorcycle/:id/sell", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.StatusDealerMotorcycle)
		r.Post("/dealer/motorcycle/:id/dont-sell", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.StatusDealerMotorcycle)
		r.Post("/dealer/motorcycle/:id", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.UpdateDealerMotorcycle)
		r.Delete("/dealer/motorcycle/:id", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.DeleteDealerMotorcycle)
		r.Post("/dealer/motorcycle/:id/images", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.CreateDealerMotorcycleImages)
		r.Post("/dealer/motorcycle/:id/videos", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.CreateDealerMotorcycleVideos)
		r.Delete("/dealer/motorcycle/:id/images", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.DeleteDealerMotorcycleImage)
		r.Delete("/dealer/motorcycle/:id/videos", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.DeleteDealerMotorcycleVideo)

		// dealer comtrans routes
		r.Post("/dealer/comtrans", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.CreateDealerComtrans)
		r.Get("/dealer/comtrans/:id/edit", auth.TokenGuard, auth.DealerGuard, auth.LanguageChecker, thirdPartyHandler.GetEditDealerComtransByID)
		r.Post("/dealer/comtrans/:id/sell", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.StatusDealerComtrans)
		r.Post("/dealer/comtrans/:id/dont-sell", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.StatusDealerComtrans)
		r.Post("/dealer/comtrans/:id", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.UpdateDealerComtrans)
		r.Delete("/dealer/comtrans/:id", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.DeleteDealerComtrans)
		r.Post("/dealer/comtrans/:id/images", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.CreateDealerComtransImages)
		r.Post("/dealer/comtrans/:id/videos", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.CreateDealerComtransVideos)
		r.Delete("/dealer/comtrans/:id/images", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.DeleteDealerComtransImage)
		r.Delete("/dealer/comtrans/:id/videos", auth.TokenGuard, auth.DealerGuard, thirdPartyHandler.DeleteDealerComtransVideo)

		// logist routes
		r.Get("/logist/destinations", auth.TokenGuard, auth.LogistGuard, auth.LanguageChecker, thirdPartyHandler.GetLogistDestinations)
		r.Post("/logist/destinations", auth.TokenGuard, auth.LogistGuard, thirdPartyHandler.CreateLogistDestination)
		r.Delete("/logist/destinations/:id", auth.TokenGuard, auth.LogistGuard, thirdPartyHandler.DeleteLogistDestination)
		// broker routes
		// car service routes
	}
}
