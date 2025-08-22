package route

import (
	_ "dubai-auto/docs"
	"dubai-auto/internal/delivery/http"
	"dubai-auto/internal/repository"
	"dubai-auto/internal/service"
	"dubai-auto/pkg/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Init(r *fiber.App, db *pgxpool.Pool) {

	userRoute := r.Group("/api/v1/users")
	SetupUserRoutes(userRoute, db)

	authRoute := r.Group("/api/v1/auth")
	SetupAuthRoutes(authRoute, db)

	motorcycleRoute := r.Group("/api/v1/motorcycles")
	SetupMotorcycleRoutes(motorcycleRoute, db)

}

func SetupUserRoutes(r fiber.Router, db *pgxpool.Pool) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := http.NewUserHandler(userService)

	{
		// profile
		r.Get("/profile/my-cars/:id", auth.TokenGuard, userHandler.GetMyCars)
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
		r.Put("/cars", auth.TokenGuard, userHandler.UpdateCar)
		r.Post("/cars/:id/images", auth.TokenGuard, userHandler.CreateCarImages)
		r.Post("/cars/:id/videos", auth.TokenGuard, userHandler.CreateCarVideos)
		r.Post("/cars/:id/cancel", auth.TokenGuard, userHandler.Cancel)
		r.Post("/cars/:id/dont-sell", auth.TokenGuard, userHandler.DontSell)
		r.Post("/cars/:id/sell", auth.TokenGuard, userHandler.Sell)
		r.Delete("/cars/:id/images", auth.TokenGuard, userHandler.DeleteCarImage)
		r.Delete("/cars/:id/videos", auth.TokenGuard, userHandler.DeleteCarVideo)
		r.Delete("/cars/:id", auth.TokenGuard, userHandler.DeleteCar)

		// likes
		r.Get("/likes", auth.TokenGuard, userHandler.Likes)
		r.Post("/likes/:car_id", auth.TokenGuard, userHandler.CarLike)
		r.Delete("/likes/:car_id", auth.TokenGuard, userHandler.RemoveLike)

	}

}

func SetupAuthRoutes(r fiber.Router, db *pgxpool.Pool) {
	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository)
	authHandler := http.NewAuthHandler(authService)

	{
		r.Post("/user-login-email", authHandler.UserLoginEmail)
		r.Post("/user-email-confirmation", authHandler.UserEmailConfirmation)
		r.Post("/user-login-phone", authHandler.UserLoginPhone)
		r.Post("/user-phone-confirmation", authHandler.UserPhoneConfirmation)
		r.Delete("/account/:id", auth.TokenGuard, authHandler.DeleteAccount)
	}
}

func SetupMotorcycleRoutes(r fiber.Router, db *pgxpool.Pool) {
	motorcycleRepository := repository.NewMotorcycleRepository(db)
	motorcycleService := service.NewMotorcycleService(motorcycleRepository)
	motorcycleHandler := http.NewMotorcycleHandler(motorcycleService)

	{
		// get motorcycles categories
		r.Get("/categories", auth.TokenGuard, motorcycleHandler.GetMotorcycleCategories)
		r.Get("/categories/:category_id/parameters", auth.TokenGuard, motorcycleHandler.GetMotorcycleParameters)
		r.Get("/categories/:category_id/brands", auth.TokenGuard, motorcycleHandler.GetMotorcycleBrands)
		r.Get("/categories/:category_id/brands/:brand_id/models", auth.TokenGuard, motorcycleHandler.GetMotorcycleModelsByBrandID)

		// motorcycles
		r.Get("/", auth.TokenGuard, motorcycleHandler.GetMotorcycles)
		r.Post("/", auth.TokenGuard, motorcycleHandler.CreateMotorcycle)
		// r.Get("/:id", auth.TokenGuard, motorcycleHandler.GetMotorcycleByID)
		// r.Get("/:id/edit", auth.TokenGuard, motorcycleHandler.GetEditMotorcycleByID)
		// r.Post("/:id/buy", auth.TokenGuard, motorcycleHandler.BuyMotorcycle)
		// r.Post("/:id/dont-sell", auth.TokenGuard, motorcycleHandler.DontSellMotorcycle)
		// r.Post("/:id/sell", auth.TokenGuard, motorcycleHandler.SellMotorcycle)
		// r.Delete("/:id", auth.TokenGuard, motorcycleHandler.DeleteMotorcycle)
		// r.Delete("/:id/images", auth.TokenGuard, motorcycleHandler.DeleteMotorcycleImage)
		// r.Delete("/:id/videos", auth.TokenGuard, motorcycleHandler.DeleteMotorcycleVideo)
	}
}
