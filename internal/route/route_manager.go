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

}

func SetupUserRoutes(r fiber.Router, db *pgxpool.Pool) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := http.NewUserHandler(userService)

	{
		r.Get("/profile/my-cars/:id", auth.TokenGuard, userHandler.GetMyCars)
		r.Get("/profile/my-cars", auth.TokenGuard, userHandler.GetMyCars)
		r.Get("/profile/on-sale", auth.TokenGuard, userHandler.OnSale)

		r.Get("/profile", auth.TokenGuard, userHandler.GetProfile)
		r.Put("/profile", auth.TokenGuard, userHandler.UpdateProfile)
		r.Get("/brands", userHandler.GetBrands)
		r.Get("/filter-brands", userHandler.GetFilterBrands)
		r.Get("/cities", userHandler.GetCities)
		r.Get("/brands/:id/models", userHandler.GetModelsByBrandID)
		r.Get("/brands/:id/filter-models", userHandler.GetFilterModelsByBrandID)
		// r.Get("/brands/filter-models", userHandler.GetFilterModelsByBrands)
		// r.Get("/brands/models/years", userHandler.GetYearsByModels)
		r.Get("/brands/:id/models/:model_id/years", userHandler.GetYearsByModelID)
		r.Get("/brands/:id/models/:model_id/body-types", userHandler.GetBodyTypesByModelID)
		// r.Get("/brands/models/body-types", userHandler.GetBodyTypesByModels)
		r.Get("/brands/:id/models/:model_id/generations", userHandler.GetGenerationsByModelID)
		r.Get("/models/generations", userHandler.GetGenerationsByModels)
		r.Get("/body-types", userHandler.GetBodyTypes)
		r.Get("/transmissions", userHandler.GetTransmissions)
		r.Get("/engines", userHandler.GetEngines)
		r.Get("/drivetrains", userHandler.GetDrivetrains)
		r.Get("/fuel-types", userHandler.GetFuelTypes)
		r.Get("/colors", userHandler.GetColors)
		r.Get("/home", auth.UserGuardOrDefault, userHandler.GetHome)
		r.Get("/cars", auth.UserGuardOrDefault, userHandler.GetCars)
		r.Get("/cars/:id", auth.UserGuardOrDefault, userHandler.GetCarByID)
		r.Get("/cars/:id/edit", auth.UserGuardOrDefault, userHandler.GetEditCarByID)
		r.Get("/likes", auth.TokenGuard, userHandler.Likes)

		r.Post("/cars/:id/buy", auth.TokenGuard, userHandler.BuyCar)
		r.Post("/cars", auth.TokenGuard, userHandler.CreateCar)
		r.Put("/cars", auth.TokenGuard, userHandler.UpdateCar)
		r.Post("/likes/:car_id", auth.TokenGuard, userHandler.CarLike)
		r.Post("/cars/:id/images", auth.TokenGuard, userHandler.CreateCarImages)
		r.Post("/cars/:id/videos", auth.TokenGuard, userHandler.CreateCarVideos)
		r.Post("/cars/:id/cancel", auth.TokenGuard, userHandler.Cancel)
		r.Post("/cars/:id/dont-sell", auth.TokenGuard, userHandler.DontSell)
		r.Post("/cars/:id/sell", auth.TokenGuard, userHandler.Sell)

		r.Delete("/likes/:car_id", auth.TokenGuard, userHandler.RemoveLike)
		r.Delete("/cars/:id/images", auth.TokenGuard, userHandler.DeleteCarImage)
		r.Delete("/cars/:id/videos", auth.TokenGuard, userHandler.DeleteCarVideo)
		r.Delete("/cars/:id", auth.TokenGuard, userHandler.DeleteCar)
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
