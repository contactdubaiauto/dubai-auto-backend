package route

import (
	"dubai-auto/internal/delivery/http"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupWebSocketRoutes(app *fiber.App, db *pgxpool.Pool) {
	// socketRepository := repository.NewSocketRepository(db)
	// socketService := service.NewSocketService(socketRepository)
	// socketHandler := http.NewSocketHandler(socketService)

	app.Use("/ws", func(c *fiber.Ctx) error {

		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}

		return fiber.ErrUpgradeRequired
	})

	http.SetupWebSocket(app)

}
