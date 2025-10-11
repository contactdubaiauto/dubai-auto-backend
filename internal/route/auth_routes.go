package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"

	"dubai-auto/internal/delivery/http"
	"dubai-auto/internal/repository"
	"dubai-auto/internal/service"
	"dubai-auto/pkg/auth"
)

func SetupAuthRoutes(r fiber.Router, db *pgxpool.Pool) {
	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository)
	authHandler := http.NewAuthHandler(authService)

	{
		r.Post("/admin-login", authHandler.AdminLogin)
		r.Post("/send-application", authHandler.Application)
		r.Post("/send-application-document", auth.TokenGuard, authHandler.ApplicationDocuments)
		r.Post("/user-login-google", authHandler.UserLoginGoogle)
		r.Post("/user-login-email", authHandler.UserLoginEmail)
		r.Post("/third-party-login", authHandler.ThirdPartyLogin)
		r.Post("/user-email-confirmation", authHandler.UserEmailConfirmation)
		r.Post("/user-login-phone", authHandler.UserLoginPhone)
		r.Post("/user-phone-confirmation", authHandler.UserPhoneConfirmation)
		r.Delete("/account/:id", auth.TokenGuard, authHandler.DeleteAccount)
	}

}
