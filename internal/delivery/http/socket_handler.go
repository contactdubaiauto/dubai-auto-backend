package http

import (
	"dubai-auto/internal/model"
	"dubai-auto/internal/service"
	"dubai-auto/internal/utils"
	"dubai-auto/pkg/auth"
	"log"
	"sync"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type SocketHandler struct {
	service *service.SocketService
}

func NewSocketHandler(service *service.SocketService) *SocketHandler {
	return &SocketHandler{service}
}

var (
	wsClients   = make(map[*websocket.Conn]*model.WSUser) // connection -> user
	wsUserConns = make(map[int][]*websocket.Conn)         // userID -> []connections
	wsMutex     = sync.RWMutex{}
)

//	{
//	    "event": "private_message",
//	    "target_user_id": 1,
//	    "data": {
//	        "message": "Hello, world!",
//	        "type": "text",
//	        "time": "2025-01-01 12:00:00"
//	    }
//	}

func (h *SocketHandler) SetupWebSocket(app *fiber.App) {

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		token := c.Query("token")
		user, err := auth.ValidateWSJWT(token)

		if err != nil {
			log.Printf("❌ Authentication failed: %v", err)
			errMsg := model.WSMessage{
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
		wsUserConns[user.ID] = append(wsUserConns[user.ID], c)
		wsMutex.Unlock()
		log.Printf("✅ User %d connected via WebSocket", user.ID)

		welcomeMsg := model.WSMessage{
			Event: "connected",
			Data: map[string]any{
				"message": "Successfully connected to messaging service",
				"user_id": user.ID,
			},
		}

		user.Avatar = h.service.GetUserAvatar(user.ID)
		wsClients[c] = user

		c.WriteJSON(welcomeMsg)
		err = h.service.UpdateUserStatus(user.ID, true)

		if err != nil {
			log.Printf("❌ Error updating user status: %v", err)
		}

		defer func() {
			wsMutex.Lock()

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
			err = h.service.UpdateUserStatus(user.ID, false)

			if err != nil {
				log.Printf("❌ Error updating user status: %v", err)
			}
			c.Close()
		}()

		data, err := h.service.GetNewMessages(user.ID)

		if err != nil {
			log.Printf("❌ Error getting unread messages: %v", err)
		} else if data != nil {
			sendToUser(user.ID, "new_message", data)
		}

		for {
			var msg model.WSMessage

			if err := c.ReadJSON(&msg); err != nil {
				log.Printf("❌ Read error: %v", err)
				break
			}

			log.Printf("📨 Received message from user %d: %s", user.ID, msg.Event)

			switch msg.Event {
			case "ping":
				pongMsg := model.WSMessage{
					Event: "pong",
					Data: map[string]any{
						"timestamp": time.Now(),
					},
				}
				c.WriteJSON(pongMsg)

			case "private_message":
				handlePrivateMessage(user, msg)

			default:
				log.Printf("⚠️ Unknown event: %s", msg.Event)
			}
		}
	}))

	log.Println("🔌 WebSocket messaging service initialized at /ws")
	log.Println("📖 Connect with: ws://localhost:8080/ws?token=YOUR_JWT_TOKEN")
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

	msg := model.WSMessage{
		Event:        event,
		TargetUserID: userID,
		Data:         data,
	}

	for _, conn := range connections {
		go func(c *websocket.Conn) {
			if err := c.WriteJSON(msg); err != nil {
				log.Printf("❌ Send error: %v", err)
			}
		}(conn)
	}
}

func handlePrivateMessage(sender *model.WSUser, msg model.WSMessage) {
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

	dataObj, exists := data["data"]

	if !exists {
		log.Printf("❌ Missing message in private message")
		return
	}

	messageData, ok := dataObj.(map[string]any)

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
		"server_time":   utils.GMTTime(),
	}

	sendToUser(int(targetUserID), "new_message", privateMessageData)

	log.Printf("📤 Private message sent from user %d to user %d (type: %v)", sender.ID, int(targetUserID), messageType)
}
