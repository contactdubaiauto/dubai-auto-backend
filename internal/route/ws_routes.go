package route

import (
	"dubai-auto/internal/delivery/http"
	"dubai-auto/internal/repository"
	"dubai-auto/internal/service"
	"dubai-auto/pkg/firebase"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupWebSocketRoutes(app *fiber.App, db *pgxpool.Pool, firebaseService *firebase.FirebaseService) {
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	socketRepository := repository.NewSocketRepository(db)
	socketService := service.NewSocketService(socketRepository, firebaseService)
	socketHandler := http.NewSocketHandler(socketService)

	socketHandler.SetupWebSocket(app)

}
