package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"

	"dubai-auto/internal/config"
	"dubai-auto/internal/delivery/http"
	"dubai-auto/internal/repository"
	"dubai-auto/internal/service"
	"dubai-auto/pkg/auth"
)

func SetupUserRoutes(r fiber.Router, config *config.Config, db *pgxpool.Pool, validator *auth.Validator) {
	userRepository := repository.NewUserRepository(config, db)
	userService := service.NewUserService(userRepository)

	motorcycleRepository := repository.NewMotorcycleRepository(config, db)
	motorcycleService := service.NewMotorcycleService(motorcycleRepository)

	comtransRepository := repository.NewComtransRepository(config, db)
	comtransService := service.NewComtransService(comtransRepository)

	userHandler := http.NewUserHandler(userService, motorcycleService, comtransService, validator)

	{
		// countries
		r.Get("/countries", auth.LanguageChecker, userHandler.GetCountries)

		// filter
		r.Get("/filter-brands", auth.LanguageChecker, userHandler.GetFilterBrands)
		r.Get("/cities", auth.LanguageChecker, userHandler.GetCities)
		r.Get("/models/generations", auth.LanguageChecker, userHandler.GetGenerationsByModels)
		r.Get("/body-types", auth.LanguageChecker, userHandler.GetBodyTypes)
		r.Get("/transmissions", auth.LanguageChecker, userHandler.GetTransmissions)
		r.Get("/engines", auth.LanguageChecker, userHandler.GetEngines)
		r.Get("/drivetrains", auth.LanguageChecker, userHandler.GetDrivetrains)
		r.Get("/fuel-types", auth.LanguageChecker, userHandler.GetFuelTypes)
		r.Get("/colors", auth.LanguageChecker, userHandler.GetColors)
		r.Get("/home", auth.UserGuardOrDefault, auth.LanguageChecker, userHandler.GetHome)
		r.Get("/third-party", userHandler.GetThirdPartyUsers)

		// profile
		r.Get("/profile/my-cars", auth.TokenGuard, auth.LanguageChecker, userHandler.GetMyCars)
		r.Get("/profile/on-sale", auth.TokenGuard, auth.LanguageChecker, userHandler.OnSale)
		r.Get("/profile", auth.TokenGuard, auth.LanguageChecker, userHandler.GetProfile)
		r.Put("/profile", auth.TokenGuard, userHandler.UpdateProfile)

		// brands
		r.Get("/brands", auth.LanguageChecker, userHandler.GetBrands)
		r.Get("/brands/:id/models", auth.LanguageChecker, userHandler.GetModelsByBrandID)
		r.Get("/brands/:id/filter-models", auth.LanguageChecker, userHandler.GetFilterModelsByBrandID)
		r.Get("/brands/:id/models/:model_id/years", userHandler.GetYearsByModelID)
		r.Get("/brands/:id/models/:model_id/body-types", auth.LanguageChecker, userHandler.GetBodyTypesByModelID)
		r.Get("/brands/:id/models/:model_id/generations", auth.LanguageChecker, userHandler.GetGenerationsByModelID)

		// cars
		r.Get("/cars", auth.UserGuardOrDefault, auth.LanguageChecker, userHandler.GetCars)
		r.Get("/cars/price-recommendation", auth.TokenGuard, userHandler.GetPriceRecommendation)
		r.Get("/cars/:id", auth.UserGuardOrDefault, auth.LanguageChecker, userHandler.GetCarByID)
		r.Get("/cars/:id/edit", auth.UserGuardOrDefault, auth.LanguageChecker, userHandler.GetEditCarByID)
		r.Post("/cars/:id/buy", auth.TokenGuard, userHandler.BuyCar)
		r.Post("/cars", auth.TokenGuard, userHandler.CreateCar)
		r.Post("/cars/:id/images", auth.TokenGuard, userHandler.CreateCarImages)
		r.Post("/cars/:id/videos", auth.TokenGuard, userHandler.CreateCarVideos)
		r.Post("/cars/:id/cancel", auth.TokenGuard, userHandler.Cancel)
		r.Post("/cars/:id/dont-sell", auth.TokenGuard, userHandler.DontSell)
		r.Post("/cars/:id/sell", auth.TokenGuard, userHandler.Sell)
		r.Put("/cars", auth.TokenGuard, userHandler.UpdateCar)
		r.Delete("/cars/:id/images", auth.TokenGuard, userHandler.DeleteCarImage)
		r.Delete("/cars/:id/videos", auth.TokenGuard, userHandler.DeleteCarVideo)
		r.Delete("/cars/:id", auth.TokenGuard, userHandler.DeleteCar)

		// motorcycles
		r.Post("/motorcycles", auth.TokenGuard, userHandler.CreateMotorcycle)
		r.Get("/motorcycles", auth.TokenGuard, userHandler.GetMotorcycles)
		r.Post("/motorcycles/:id/buy", auth.TokenGuard, userHandler.BuyMotorcycle)
		r.Post("/motorcycles/:id/cancel", auth.TokenGuard, userHandler.CancelMotorcycle)
		r.Post("/motorcycles/:id/dont-sell", auth.TokenGuard, userHandler.DontSellMotorcycle)
		r.Post("/motorcycles/:id/sell", auth.TokenGuard, userHandler.SellMotorcycle)
		r.Put("/motorcycles", auth.TokenGuard, userHandler.UpdateMotorcycle)
		r.Delete("/motorcycles/:id/images", auth.TokenGuard, userHandler.DeleteMotorcycleImage)
		r.Delete("/motorcycles/:id/videos", auth.TokenGuard, userHandler.DeleteMotorcycleVideo)
		r.Delete("/motorcycles/:id", auth.TokenGuard, userHandler.DeleteMotorcycle)

		// comtrans
		r.Post("/comtrans", auth.TokenGuard, auth.LanguageChecker, userHandler.CreateComtrans)
		r.Get("/comtrans", auth.TokenGuard, auth.LanguageChecker, userHandler.GetComtrans)
		r.Post("/comtrans/:id/buy", auth.TokenGuard, auth.LanguageChecker, userHandler.BuyComtrans)
		r.Post("/comtrans/:id/cancel", auth.TokenGuard, auth.LanguageChecker, userHandler.CancelComtrans)
		r.Post("/comtrans/:id/dont-sell", auth.TokenGuard, auth.LanguageChecker, userHandler.DontSellComtrans)
		r.Post("/comtrans/:id/sell", auth.TokenGuard, auth.LanguageChecker, userHandler.SellComtrans)
		r.Put("/comtrans", auth.TokenGuard, auth.LanguageChecker, userHandler.UpdateComtrans)
		r.Delete("/comtrans/:id/images", auth.TokenGuard, auth.LanguageChecker, userHandler.DeleteComtransImage)
		r.Delete("/comtrans/:id/videos", auth.TokenGuard, auth.LanguageChecker, userHandler.DeleteComtransVideo)
		r.Delete("/comtrans/:id", auth.TokenGuard, auth.LanguageChecker, userHandler.DeleteComtrans)

		// likes
		r.Get("/likes", auth.TokenGuard, auth.LanguageChecker, userHandler.Likes)
		r.Post("/likes/:car_id", auth.TokenGuard, userHandler.CarLike)
		r.Delete("/likes/:car_id", auth.TokenGuard, userHandler.RemoveLike)

		// messages
		r.Post("/messages/files", auth.TokenGuard, userHandler.CreateMessageFile)

		// users
		r.Get("/:id", auth.LanguageChecker, userHandler.GetUserByID)

		// Report
		r.Post("/reports", auth.TokenGuard, userHandler.CreateReport)
		r.Get("/reports", auth.TokenGuard, userHandler.GetReports)
		r.Post("/item-reports", auth.TokenGuard, userHandler.CreateItemReports)
		r.Get("/item-reports", auth.TokenGuard, userHandler.GetItemReports)
	}
}
