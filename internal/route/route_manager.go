package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"dubai-auto/internal/delivery/http"
	"dubai-auto/internal/repository"
	"dubai-auto/internal/service"
	"dubai-auto/pkg"
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
		r.GET("/profile/cars", pkg.TokenGuard, userHandler.GetProfileCars)

		r.GET("/brands", userHandler.GetBrands)
		r.GET("/brands/:id/models", userHandler.GetModelsByBrandID)
		r.GET("/brands/:id/models/:model_id/generations", userHandler.GetGenerationsByModelID)
		r.GET("/body-types", userHandler.GetBodyTypes)
		r.GET("/transmissions", userHandler.GetTransmissions)
		r.GET("/engines", userHandler.GetEngines)
		r.GET("/drivetrains", userHandler.GetDrivetrains)
		r.GET("/fuel-types", userHandler.GetFuelTypes)
		r.GET("/colors", userHandler.GetColors)
		r.GET("/cars", userHandler.GetCars)
		r.GET("/cars/:id", userHandler.GetCarByID)
		r.POST("/cars", pkg.TokenGuard, userHandler.CreateCar)
		r.POST("/cars/:id/images", pkg.TokenGuard, userHandler.CreateCarImages)

	}
}

func SetupAuthRoutes(r *gin.RouterGroup, db *pgxpool.Pool) {
	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository)
	authHandler := http.NewAuthHandler(authService)

	{
		r.POST("/user-login", authHandler.UserLogin)
		r.POST("/user-register", authHandler.UserRegister)
	}
}
