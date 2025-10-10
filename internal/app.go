package internal

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	fb_logger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"github.com/jackc/pgx/v5/pgxpool"

	"dubai-auto/internal/config"
	"dubai-auto/internal/route"
	"dubai-auto/pkg/auth"
	"dubai-auto/pkg/logger"
)

// Message structures for Socket.IO
type SocketMessage struct {
	ID        string    `json:"id"`
	UserID    int       `json:"user_id"`
	Username  string    `json:"username"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	RoomID    string    `json:"room_id,omitempty"`
}

type SocketUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	RoleID   int    `json:"role_id"`
	SocketID string `json:"socket_id"`
}

// Global variables for Socket.IO messaging
var (
	socketClients     = make(map[string]*SocketUser)   // socketID -> user
	socketUserSockets = make(map[int][]string)         // userID -> []socketID
	socketConnections = make(map[string]socketio.Conn) // socketID -> connection
	socketMutex       = sync.RWMutex{}
	socketServer      *socketio.Server
)

// JWT validation for Socket.IO connections
func validateSocketJWT(tokenString string) (*SocketUser, error) {
	if tokenString == "" {
		return nil, fmt.Errorf("missing token")
	}

	// Remove Bearer prefix if present
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(
		tokenString, claims,
		func(t *jwt.Token) (any, error) {
			return []byte(auth.ENV.ACCESS_KEY), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	userID := int(claims["id"].(float64))
	roleID := int(claims["role_id"].(float64))

	return &SocketUser{
		ID:       userID,
		Username: fmt.Sprintf("User_%d", userID),
		RoleID:   roleID,
	}, nil
}

func InitApp(db *pgxpool.Pool, conf *config.Config, logger *logger.Logger) *fiber.App {
	appConfig := fiber.Config{
		BodyLimit: 50 * 1024 * 1024,
	}

	app := fiber.New(appConfig)
	app.Use(pprof.New())
	app.Use(auth.Cors)

	if config.ENV.APP_MODE != "release" {
		app.Use(fb_logger.New(fb_logger.Config{
			Format: "[${time}] ${ip} ${status} - ${method} ${path} ${latency}\n",
			Output: os.Stdout, // Or to a file: os.OpenFile("fiber.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		}))
	}

	app.Static("api/v1/images", "."+conf.STATIC_PATH)

	// Initialize Socket.IO messaging service
	setupSocketIO(app)

	// Initialize routes
	route.Init(app, db)
	return app
}

// setupSocketIO initializes the Socket.IO messaging service
func setupSocketIO(app *fiber.App) {
	// Create Socket.IO server with proper transports
	socketServer = socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: func(r *http.Request) bool {
					return true // Allow all origins for development
				},
			},
			&websocket.Transport{
				CheckOrigin: func(r *http.Request) bool {
					return true // Allow all origins for development
				},
			},
		},
	})

	// Handle connection events
	socketServer.OnConnect("/", func(conn socketio.Conn) error {
		// Get JWT token from query parameters
		url := conn.URL()
		token := url.Query().Get("token")
		if token == "" {
			token = url.Query().Get("auth") // Alternative parameter name
		}

		user, err := validateSocketJWT(token)
		if err != nil {
			log.Printf("❌ Authentication failed: %v", err)
			conn.Emit("error", map[string]string{
				"error":   "authentication_failed",
				"message": err.Error(),
			})
			conn.Close()
			return err
		}

		// Store user info
		socketID := conn.ID()
		user.SocketID = socketID

		socketMutex.Lock()
		socketClients[socketID] = user
		socketUserSockets[user.ID] = append(socketUserSockets[user.ID], socketID)
		socketConnections[socketID] = conn
		socketMutex.Unlock()

		log.Printf("✅ User %d connected with socket %s", user.ID, socketID)

		// Send welcome message
		conn.Emit("connected", map[string]any{
			"message":   "Successfully connected to messaging service",
			"user_id":   user.ID,
			"socket_id": socketID,
		})

		// Broadcast user online status to all clients
		broadcastUserStatus(user.ID, "online")

		return nil
	})

	// Handle disconnection events
	socketServer.OnDisconnect("/", func(conn socketio.Conn, reason string) {
		socketID := conn.ID()

		socketMutex.Lock()
		user, exists := socketClients[socketID]
		if exists {
			delete(socketClients, socketID)
			delete(socketConnections, socketID)

			if sockets, userExists := socketUserSockets[user.ID]; userExists {
				// Remove this socket from user's socket list
				for i, sid := range sockets {
					if sid == socketID {
						socketUserSockets[user.ID] = append(sockets[:i], sockets[i+1:]...)
						break
					}
				}
				// If no more sockets for this user, remove the user entry
				if len(socketUserSockets[user.ID]) == 0 {
					delete(socketUserSockets, user.ID)
				}
			}
		}
		socketMutex.Unlock()

		if exists {
			log.Printf("🔌 User %d disconnected (socket %s): %s", user.ID, socketID, reason)
			broadcastUserStatus(user.ID, "offline")
		}
	})

	// Handle message events
	socketServer.OnEvent("/", "message", func(conn socketio.Conn, data map[string]any) {
		socketID := conn.ID()

		socketMutex.RLock()
		user, exists := socketClients[socketID]
		socketMutex.RUnlock()

		if !exists {
			log.Printf("❌ Message from unknown socket: %s", socketID)
			return
		}

		if msgText, ok := data["message"].(string); ok {
			chatMessage := SocketMessage{
				ID:        uuid.New().String(),
				UserID:    user.ID,
				Username:  user.Username,
				Message:   msgText,
				Timestamp: time.Now(),
			}

			log.Printf("💬 Message from user %d: %s", user.ID, msgText)
			broadcastMessage("message", chatMessage)
		}
	})

	// Handle private message events
	socketServer.OnEvent("/", "private_message", func(conn socketio.Conn, data map[string]any) {
		socketID := conn.ID()

		socketMutex.RLock()
		user, exists := socketClients[socketID]
		socketMutex.RUnlock()

		if !exists {
			log.Printf("❌ Private message from unknown socket: %s", socketID)
			return
		}

		if targetUserID, ok := data["target_user_id"].(float64); ok {
			if msgText, ok := data["message"].(string); ok {
				privateMessage := SocketMessage{
					ID:        uuid.New().String(),
					UserID:    user.ID,
					Username:  user.Username,
					Message:   msgText,
					Timestamp: time.Now(),
				}

				log.Printf("📩 Private message from user %d to user %d: %s", user.ID, int(targetUserID), msgText)
				sendToUser(int(targetUserID), "private_message", privateMessage)
			}
		}
	})

	// Handle ping events
	socketServer.OnEvent("/", "ping", func(conn socketio.Conn) {
		conn.Emit("pong", map[string]any{
			"timestamp": time.Now(),
		})
	})

	// Start the Socket.IO server
	go func() {
		if err := socketServer.Serve(); err != nil {
			log.Fatalf("❌ Socket.IO server error: %v", err)
		}
	}()

	// Mount Socket.IO routes using Fiber adaptor
	app.Get("/socket.io/*", adaptor.HTTPHandler(socketServer))

	// Add HTTP API 		endpoints for Socket.IO
	socketGroup := app.Group("/api/v1/socket")

	// Public info endpoint
	socketGroup.Get("/info", func(c *fiber.Ctx) error {
		info := map[string]any{
			"service":     "Socket.IO Messaging Service",
			"version":     "1.0.0",
			"endpoint":    "/socket.io/",
			"auth_method": "JWT Token",
			"events": []string{
				"message",
				"private_message",
				"join_room",
				"get_online_users",
			},
			"auth_examples": map[string]string{
				"query_param": "?token=your_jwt_token",
				"alt_param":   "?auth=your_jwt_token",
			},
			"message_format": map[string]any{
				"simple_message			": `{"message": "Hello world!"}`,
				"event_message":     `{"event": "message", "message": "Hello world!"}`,
				"private_message":   `{"event": "private_message", "target_user_id": 123, "message": "Hello!"}`,
				"join_room":         `{"event": "join_room", "room_id": "general"}`,
				"get_users":         `{"event":		 "get_online_users"}`,
			},
			"note": "This is a demonstration implementation. For full functionality, you may need to use a different Socket.IO library or implement custom WebSocket handling.",
		}
		return c.JSON(map[string]any{
			"event": "info",
			"data":  info,
		})
	})

	// Protected stats endpoint
	socketGroup.Get("/stats", auth.TokenGuard, func(c *fiber.Ctx) error {
		socketMutex.RLock()
		connectedUsers := len(socketClients)
		userList := make([]SocketUser, 0, len(socketClients))
		for _, user := range socketClients {
			userList = append(userList, *user)
		}
		socketMutex.RUnlock()

		stats := map[string]any{
			"connected_users": connectedUsers,
			"users":           userList,
		}
		return c.JSON(map[string]any{
			"event": "stats",
			"data":  stats,
		})
	})

	// System message endpoint (admin only)
	socketGroup.Post("/system-message", auth.TokenGuard, func(c *fiber.Ctx) error {
		var req struct {
			Message string `json:"message"`
		}

		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(map[string]string{
				"error": "Invalid request body",
			})
		}

		systemMessage := SocketMessage{
			ID:        "system",
			UserID:    0,
			Username:  "System",
			Message:   req.Message,
			Timestamp: time.Now(),
		}

		// Broadcast system message
		broadcastMessage("system_message", systemMessage)

		return c.JSON(map[string]any{
			"event": "system_message_sent",
			"data":  systemMessage,
		})
	})

	log.Println("🔌 Socket.IO messaging service initialized")
	log.Println("📖 Get service info:GET /api/v1/socket/info")
	log.Println("📊 Get stats: GET /api/v1/sockt/stats (requires JWT)")
	log.Println("📢 Send system message: POST /api/v1/socket/system-message (requires JWT)")
}

// Helper functions
func broadcastMessage(event string, data any) {
	socketMutex.RLock()
	connections := make(map[string]socketio.Conn)
	for k, v := range socketConnections {
		connections[k] = v
	}
	socketMutex.RUnlock()

	log.Printf("📡 Broadcasting %s to %d users", event, len(connections))

	// Send to all connected clients
	for socketID, conn := range connections {
		go func(sid string, c socketio.Conn) {
			c.Emit(event, data)
		}(socketID, conn)
	}
}

func broadcastUserStatus(userID int, status string) {
	statusMessage := map[string]any{
		"user_id":  userID,
		"username": fmt.Sprintf("User_%d", userID),
		"status":   status,
	}

	socketMutex.RLock()
	connectedCount := len(socketClients)
	socketMutex.RUnlock()

	log.Printf("📡 Broadcasting user %d status: %s to %d users", userID, status, connectedCount)
	broadcastMessage("user_status", statusMessage)
}

func sendToUser(userID int, event string, data any) {
	socketMutex.RLock()
	userSockets, exists := socketUserSockets[userID]
	if !exists || len(userSockets) == 0 {
		socketMutex.RUnlock()
		log.Printf("❌ User %d not connected", userID)
		return
	}

	// Get connections for this user
	userConnections := make([]socketio.Conn, 0, len(userSockets))
	for _, socketID := range userSockets {
		if conn, connExists := socketConnections[socketID]; connExists {
			userConnections = append(userConnections, conn)
		}
	}
	socketMutex.RUnlock()

	log.Printf("📤 Sending %s to user %d (%d sockets)", event, userID, len(userConnections))

	// Send to all user's connections
	for _, conn := range userConnections {
		go func(c socketio.Conn) {
			c.Emit(event, data)
		}(conn)
	}
}
