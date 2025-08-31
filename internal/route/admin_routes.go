package route

import (
	"dubai-auto/internal/delivery/http"
	"dubai-auto/internal/repository"
	"dubai-auto/internal/service"
	"dubai-auto/pkg/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupAdminRoutes(r fiber.Router, db *pgxpool.Pool) {
	adminRepository := repository.NewAdminRepository(db)
	adminService := service.NewAdminService(adminRepository)
	adminHandler := http.NewAdminHandler(adminService)

	// Cities routes
	cities := r.Group("/cities")
	{
		cities.Get("/", auth.TokenGuard, adminHandler.GetCities)
		cities.Post("/", auth.TokenGuard, adminHandler.CreateCity)
		cities.Put("/:id", auth.TokenGuard, adminHandler.UpdateCity)
		cities.Delete("/:id", auth.TokenGuard, adminHandler.DeleteCity)
	}

	// Brands routes
	brands := r.Group("/brands")
	{
		brands.Get("/", auth.TokenGuard, adminHandler.GetBrands)
		brands.Post("/", auth.TokenGuard, adminHandler.CreateBrand)
		brands.Put("/:id", auth.TokenGuard, adminHandler.UpdateBrand)
		brands.Delete("/:id", auth.TokenGuard, adminHandler.DeleteBrand)
	}

	// Models routes
	models := r.Group("/models")
	{
		models.Get("/", auth.TokenGuard, adminHandler.GetModels)
		models.Post("/", auth.TokenGuard, adminHandler.CreateModel)
		models.Put("/:id", auth.TokenGuard, adminHandler.UpdateModel)
		models.Delete("/:id", auth.TokenGuard, adminHandler.DeleteModel)
	}

	// Body Types routes
	bodyTypes := r.Group("/body-types")
	{
		bodyTypes.Get("/", auth.TokenGuard, adminHandler.GetBodyTypes)
		bodyTypes.Post("/", auth.TokenGuard, adminHandler.CreateBodyType)
		bodyTypes.Put("/:id", auth.TokenGuard, adminHandler.UpdateBodyType)
		bodyTypes.Delete("/:id", auth.TokenGuard, adminHandler.DeleteBodyType)
	}
}
