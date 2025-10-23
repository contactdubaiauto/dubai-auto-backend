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

func SetupUserRoutes(r fiber.Router, config *config.Config, db *pgxpool.Pool) {
	userRepository := repository.NewUserRepository(config, db)
	userService := service.NewUserService(userRepository)
	userHandler := http.NewUserHandler(userService)

	{

		// countries
		r.Get("/countries", userHandler.GetCountries)

		// profile
		r.Get("/profile/my-cars/:id", auth.TokenGuard, userHandler.GetMyCars) //todo: write GetMyCarsByID
		r.Get("/profile/my-cars", auth.TokenGuard, userHandler.GetMyCars)
		r.Get("/profile/on-sale", auth.TokenGuard, userHandler.OnSale)
		r.Get("/profile", auth.TokenGuard, userHandler.GetProfile)
		r.Put("/profile", auth.TokenGuard, userHandler.UpdateProfile)

		// brands
		r.Get("/brands", userHandler.GetBrands)
		r.Get("/brands/:id/models", userHandler.GetModelsByBrandID)
		r.Get("/brands/:id/filter-models", userHandler.GetFilterModelsByBrandID)
		r.Get("/brands/:id/models/:model_id/years", userHandler.GetYearsByModelID)
		r.Get("/brands/:id/models/:model_id/body-types", userHandler.GetBodyTypesByModelID)
		r.Get("/brands/:id/models/:model_id/generations", userHandler.GetGenerationsByModelID)

		// filter
		r.Get("/filter-brands", userHandler.GetFilterBrands)
		r.Get("/cities", userHandler.GetCities)
		// r.Get("/brands/filter-models", userHandler.GetFilterModelsByBrands)
		// r.Get("/brands/models/years", userHandler.GetYearsByModels)
		// r.Get("/brands/models/body-types", userHandler.GetBodyTypesByModels)
		r.Get("/models/generations", userHandler.GetGenerationsByModels)
		r.Get("/body-types", userHandler.GetBodyTypes)
		r.Get("/transmissions", userHandler.GetTransmissions)
		r.Get("/engines", userHandler.GetEngines)
		r.Get("/drivetrains", userHandler.GetDrivetrains)
		r.Get("/fuel-types", userHandler.GetFuelTypes)
		r.Get("/colors", userHandler.GetColors)
		r.Get("/home", auth.UserGuardOrDefault, userHandler.GetHome)

		// cars
		r.Get("/cars", auth.UserGuardOrDefault, userHandler.GetCars)
		r.Get("/cars/price-recommendation", auth.TokenGuard, userHandler.GetPriceRecommendation)
		r.Get("/cars/:id", auth.UserGuardOrDefault, userHandler.GetCarByID)
		r.Get("/cars/:id/edit", auth.UserGuardOrDefault, userHandler.GetEditCarByID)
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

		// likes
		r.Get("/likes", auth.TokenGuard, userHandler.Likes)
		r.Post("/likes/:car_id", auth.TokenGuard, userHandler.CarLike)
		r.Delete("/likes/:car_id", auth.TokenGuard, userHandler.RemoveLike)

		// brokers
		r.Get("/brokers", userHandler.GetBrokers)
		r.Get("/brokers/:id", userHandler.GetBrokerByID)

		// logists
		r.Get("/logists", userHandler.GetLogists)
		r.Get("/logists/:id", userHandler.GetLogistByID)

		// services
		r.Get("/services", userHandler.GetServices)
		r.Get("/services/:id", userHandler.GetServiceByID)
	}

}
