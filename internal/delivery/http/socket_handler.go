package http

import (
	"dubai-auto/internal/service"
	"dubai-auto/pkg/auth"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type SocketHandler struct {
	service *service.SocketService
}

func NewSocketHandler(service *service.SocketService) *SocketHandler {
	return &SocketHandler{service}
}

type WSMessage struct {
	Event   string `json:"event"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type WSUser struct {
	ID       int             `json:"id"`
	Username string          `json:"username"`
	RoleID   int             `json:"role_id"`
	Conn     *websocket.Conn `json:"-"`
}

var (
	wsClients   = make(map[*websocket.Conn]*WSUser) // connection -> user
	wsUserConns = make(map[int][]*websocket.Conn)   // userID -> []connections
	wsMutex     = sync.RWMutex{}
)

func validateWSJWT(tokenString string) (*WSUser, error) {
	if tokenString == "" {
		return nil, fmt.Errorf("missing token")
	}

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

	return &WSUser{
		ID:       userID,
		Username: fmt.Sprintf("User_%d", userID),
		RoleID:   roleID,
	}, nil
}

func SetupWebSocket(app *fiber.App) {

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		token := c.Query("token")

		if token == "" {
			token = c.Query("auth")
		}

		user, err := validateWSJWT(token)

		if err != nil {
			log.Printf("❌ Authentication failed: %v", err)
			errMsg := WSMessage{
				Event: "error",
				Data: map[string]string{
					"error":   "authentication_failed",
					"message": err.Error(),
				},
			}
			c.WriteJSON(errMsg)
			c.Close()
			return
		}

		user.Conn = c

		wsMutex.Lock()
		wsClients[c] = user
		wsUserConns[user.ID] = append(wsUserConns[user.ID], c)
		wsMutex.Unlock()

		log.Printf("✅ User %d connected via WebSocket", user.ID)

		welcomeMsg := WSMessage{
			Event: "connected",
			Data: map[string]any{
				"message": "Successfully connected to messaging service",
				"user_id": user.ID,
			},
		}
		c.WriteJSON(welcomeMsg)

		broadcastUserStatus(user.ID, "online")

		defer func() {
			wsMutex.Lock()
			delete(wsClients, c)

			if conns, exists := wsUserConns[user.ID]; exists {

				for i, conn := range conns {
					if conn == c {
						wsUserConns[user.ID] = append(conns[:i], conns[i+1:]...)
						break
					}
				}

				if len(wsUserConns[user.ID]) == 0 {
					delete(wsUserConns, user.ID)
				}
			}

			wsMutex.Unlock()
			log.Printf("🔌 User %d disconnected", user.ID)
			broadcastUserStatus(user.ID, "offline")
			c.Close()
		}()

		for {
			var msg WSMessage
			if err := c.ReadJSON(&msg); err != nil {
				log.Printf("❌ Read error: %v", err)
				break
			}

			log.Printf("📨 Received message from user %d: %s", user.ID, msg.Event)

			switch msg.Event {
			case "ping":
				pongMsg := WSMessage{
					Event: "pong",
					Data: map[string]any{
						"timestamp": time.Now(),
					},
				}
				c.WriteJSON(pongMsg)

			case "message":
				handleMessage(user, msg.Message)

			case "private_message":
				handlePrivateMessage(user, msg)

			case "get_online_users":
				sendOnlineUsers(c)

			default:
				log.Printf("⚠️ Unknown event: %s", msg.Event)
			}
		}
	}))

	log.Println("🔌 WebSocket messaging service initialized at /ws")
	log.Println("📖 Connect with: ws://localhost:8080/ws?token=YOUR_JWT_TOKEN")
}

func broadcastMessage(event string, data any) {
	wsMutex.RLock()
	connections := make([]*websocket.Conn, 0, len(wsClients))
	for conn := range wsClients {
		connections = append(connections, conn)
	}
	wsMutex.RUnlock()

	log.Printf("📡 Broadcasting %s to %d users", event, len(connections))

	msg := WSMessage{
		Event: event,
		Data:  data,
	}

	for _, conn := range connections {
		go func(c *websocket.Conn) {
			if err := c.WriteJSON(msg); err != nil {
				log.Printf("❌ Broadcast error: %v", err)
			}
		}(conn)
	}
}

func broadcastUserStatus(userID int, status string) {
	statusMessage := map[string]any{
		"user_id":  userID,
		"username": fmt.Sprintf("User_%d", userID),
		"status":   status,
	}

	wsMutex.RLock()
	connectedCount := len(wsClients)
	wsMutex.RUnlock()

	log.Printf("📡 Broadcasting user %d status: %s to %d users", userID, status, connectedCount)
	broadcastMessage("user_status", statusMessage)
}

func sendToUser(userID int, event string, data any) {
	wsMutex.RLock()
	userConns, exists := wsUserConns[userID]

	if !exists || len(userConns) == 0 {
		wsMutex.RUnlock()
		log.Printf("❌ User %d not connected", userID)
		return
	}

	connections := make([]*websocket.Conn, len(userConns))
	copy(connections, userConns)
	wsMutex.RUnlock()

	log.Printf("📤 Sending %s to user %d (%d connections)", event, userID, len(connections))

	msg := WSMessage{
		Event: event,
		Data:  data,
	}

	for _, conn := range connections {
		go func(c *websocket.Conn) {
			if err := c.WriteJSON(msg); err != nil {
				log.Printf("❌ Send error: %v", err)
			}
		}(conn)
	}
}

func handleMessage(user *WSUser, message string) {
	broadcastMessage("message", map[string]any{
		"user_id":  user.ID,
		"username": user.Username,
		"message":  message,
	})
}

func sendOnlineUsers(conn *websocket.Conn) {
	wsMutex.RLock()
	users := make([]map[string]any, 0, len(wsClients))

	for _, user := range wsClients {
		users = append(users, map[string]any{
			"id":       user.ID,
			"username": user.Username,
		})
	}
	wsMutex.RUnlock()

	msg := WSMessage{
		Event: "online_users",
		Data: map[string]any{
			"users": users,
			"count": len(users),
		},
	}

	if err := conn.WriteJSON(msg); err != nil {
		log.Printf("❌ Send online users error: %v", err)
	}
}

func handlePrivateMessage(sender *WSUser, msg WSMessage) {
	data, ok := msg.Data.(map[string]any)
	if !ok {
		log.Printf("❌ Invalid private message data format")
		return
	}

	targetUserIDFloat, exists := data["target_user_id"]
	if !exists {
		log.Printf("❌ Missing target_user_id in private message")
		return
	}

	targetUserID, ok := targetUserIDFloat.(float64)
	if !ok {
		log.Printf("❌ Invalid target_user_id type")
		return
	}

	messageObj, exists := data["message"]
	if !exists {
		log.Printf("❌ Missing message in private message")
		return
	}

	messageData, ok := messageObj.(map[string]any)
	if !ok {
		log.Printf("❌ Invalid message format - expected object with time, message, type")
		return
	}

	messageText, hasMessage := messageData["message"]
	messageType, hasType := messageData["type"]
	messageTime, hasTime := messageData["time"]

	if !hasMessage || !hasType || !hasTime {
		log.Printf("❌ Message object missing required fields: time, message, type")
		return
	}

	privateMessageData := map[string]any{
		"from_user_id":  sender.ID,
		"from_username": sender.Username,
		"message":       messageText,
		"type":          messageType,
		"time":          messageTime,
		"server_time":   time.Now(),
	}

	sendToUser(int(targetUserID), "private_message", privateMessageData)

	log.Printf("📤 Private message sent from user %d to user %d (type: %v)", sender.ID, int(targetUserID), messageType)
}
