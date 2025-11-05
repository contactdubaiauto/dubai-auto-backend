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
	userHandler := http.NewUserHandler(userService, validator)

	{

		// users
		r.Get("/:id", auth.LanguageChecker, userHandler.GetUserByID)

		// countries
		r.Get("/countries", userHandler.GetCountries)

		// filter
		r.Get("/filter-brands", auth.LanguageChecker, userHandler.GetFilterBrands)
		r.Get("/cities", userHandler.GetCities)
		// r.Get("/brands/filter-models", userHandler.GetFilterModelsByBrands)
		// r.Get("/brands/models/years", userHandler.GetYearsByModels)
		// r.Get("/brands/models/body-types", userHandler.GetBodyTypesByModels)
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
		profilesR := r.Group("/profiles")
		profilesR.Get("/profile/my-cars/:id", auth.TokenGuard, auth.LanguageChecker, userHandler.GetMyCars) //todo: write GetMyCarsByID
		profilesR.Get("/profile/my-cars", auth.TokenGuard, auth.LanguageChecker, userHandler.GetMyCars)
		profilesR.Get("/profile/on-sale", auth.TokenGuard, auth.LanguageChecker, userHandler.OnSale)
		profilesR.Get("/profile", auth.TokenGuard, userHandler.GetProfile)
		profilesR.Put("/profile", auth.TokenGuard, userHandler.UpdateProfile)

		// brands
		brandsR := r.Group("/brands")
		brandsR.Get("/brands", auth.LanguageChecker, userHandler.GetBrands)
		brandsR.Get("/brands/:id/models", auth.LanguageChecker, userHandler.GetModelsByBrandID)
		brandsR.Get("/brands/:id/filter-models", auth.LanguageChecker, userHandler.GetFilterModelsByBrandID)
		brandsR.Get("/brands/:id/models/:model_id/years", userHandler.GetYearsByModelID)
		brandsR.Get("/brands/:id/models/:model_id/body-types", auth.LanguageChecker, userHandler.GetBodyTypesByModelID)
		brandsR.Get("/brands/:id/models/:model_id/generations", auth.LanguageChecker, userHandler.GetGenerationsByModelID)

		// cars
		carsR := r.Group("/cars")
		carsR.Get("/cars", auth.UserGuardOrDefault, auth.LanguageChecker, userHandler.GetCars)
		carsR.Get("/cars/price-recommendation", auth.TokenGuard, userHandler.GetPriceRecommendation)
		carsR.Get("/cars/:id", auth.UserGuardOrDefault, auth.LanguageChecker, userHandler.GetCarByID)
		carsR.Get("/cars/:id/edit", auth.UserGuardOrDefault, auth.LanguageChecker, userHandler.GetEditCarByID)
		carsR.Post("/cars/:id/buy", auth.TokenGuard, userHandler.BuyCar)
		carsR.Post("/cars", auth.TokenGuard, userHandler.CreateCar)
		carsR.Post("/cars/:id/images", auth.TokenGuard, userHandler.CreateCarImages)
		carsR.Post("/cars/:id/videos", auth.TokenGuard, userHandler.CreateCarVideos)
		carsR.Post("/cars/:id/cancel", auth.TokenGuard, userHandler.Cancel)
		carsR.Post("/cars/:id/dont-sell", auth.TokenGuard, userHandler.DontSell)
		carsR.Post("/cars/:id/sell", auth.TokenGuard, userHandler.Sell)
		carsR.Put("/cars", auth.TokenGuard, userHandler.UpdateCar)
		carsR.Delete("/cars/:id/images", auth.TokenGuard, userHandler.DeleteCarImage)
		carsR.Delete("/cars/:id/videos", auth.TokenGuard, userHandler.DeleteCarVideo)
		carsR.Delete("/cars/:id", auth.TokenGuard, userHandler.DeleteCar)

		// likes
		likesR := r.Group("/likes")
		likesR.Get("/likes", auth.TokenGuard, auth.LanguageChecker, userHandler.Likes)
		likesR.Post("/likes/:car_id", auth.TokenGuard, userHandler.CarLike)
		likesR.Delete("/likes/:car_id", auth.TokenGuard, userHandler.RemoveLike)

		// messages
		messageR := r.Group("/messages")
		messageR.Post("/files", auth.TokenGuard, userHandler.CreateMessageFile)
	}
}
