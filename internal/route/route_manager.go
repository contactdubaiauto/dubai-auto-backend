package route

import (
	_ "dubai-auto/docs"
	"dubai-auto/internal/delivery/http"
	"dubai-auto/internal/repository"
	"dubai-auto/internal/service"
	"dubai-auto/pkg"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Init(r *gin.Engine, db *pgxpool.Pool) {

	userRoute := r.Group("/api/v1/users")
	SetupUserRoutes(userRoute, db)

	authRoute := r.Group("/api/v1/auth")
	SetupAuthRoutes(authRoute, db)

}

func SetupUserRoutes(r *gin.RouterGroup, db *pgxpool.Pool) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := http.NewUserHandler(userService)

	{
		r.GET("/profile/my-cars", pkg.TokenGuard, userHandler.GetMyCars)
		r.GET("/profile/on-sale", pkg.TokenGuard, userHandler.OnSale)

		r.GET("/brands", userHandler.GetBrands)
		r.GET("/filter-brands", userHandler.GetFilterBrands)
		r.GET("/cities", userHandler.GetCities)
		r.GET("/modifications", userHandler.GetModifications)
		r.GET("/brands/:id/models", userHandler.GetModelsByBrandID)
		r.GET("/brands/:id/filter-models", userHandler.GetFilterModelsByBrandID)
		r.GET("/brands/:id/models/:model_id/years", userHandler.GetYearsByModelID)
		r.GET("/brands/:id/models/:model_id/body-types", userHandler.GetBodyTypesByModelID)
		r.GET("/brands/:id/models/:model_id/generations", userHandler.GetGenerationsByModelID)
		r.GET("/body-types", userHandler.GetBodyTypes)
		r.GET("/transmissions", userHandler.GetTransmissions)
		r.GET("/engines", userHandler.GetEngines)
		r.GET("/drivetrains", userHandler.GetDrivetrains)
		r.GET("/fuel-types", userHandler.GetFuelTypes)
		r.GET("/colors", userHandler.GetColors)
		r.GET("/cars", pkg.UserGuardOrDefault, userHandler.GetCars)
		r.GET("/cars/:id", pkg.UserGuardOrDefault, userHandler.GetCarByID)

		r.POST("/cars/:id/buy", pkg.TokenGuard, userHandler.BuyCar)
		r.POST("/cars", pkg.TokenGuard, userHandler.CreateCar)
		r.PUT("/cars", pkg.TokenGuard, userHandler.UpdateCar)
		r.POST("/cars/:id/images", pkg.TokenGuard, userHandler.CreateCarImages)
		r.POST("/cars/:id/cancel", pkg.TokenGuard, userHandler.Cancel)
		r.POST("/cars/:id/delete", pkg.TokenGuard, userHandler.Delete)
		r.POST("/cars/:id/dont-sell", pkg.TokenGuard, userHandler.DontSell)
		r.POST("/cars/:id/sell", pkg.TokenGuard, userHandler.Sell)

	}
}

func SetupAuthRoutes(r *gin.RouterGroup, db *pgxpool.Pool) {
	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository)
	authHandler := http.NewAuthHandler(authService)

	{
		r.POST("/user-login-email", authHandler.UserLoginEmail)
		r.POST("/user-email-confirmation", authHandler.UserEmailConfirmation)
		r.POST("/user-login-phone", authHandler.UserLoginPhone)
		r.POST("/user-phone-confirmation", authHandler.UserPhoneConfirmation)
	}
}
