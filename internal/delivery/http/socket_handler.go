package http

import (
	"dubai-auto/internal/model"
	"dubai-auto/internal/service"
	"dubai-auto/internal/utils"
	"dubai-auto/pkg/auth"
	"fmt"
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
				Data: []model.UserMessage{
					{
						Messages: []model.Message{
							{
								CreatedAt: time.Now(),
								Message:   "authentication_failed",
							},
						},
					},
				},
			}

			c.WriteJSON(errMsg)
			c.Close()
			return
		}

		err = h.service.CheckUserExists(user.ID)

		if err != nil {
			log.Printf("❌ Error checking user exists: %v", err)
			errMsg := model.WSMessage{
				Event: "error",
				Data: []model.UserMessage{
					{
						Messages: []model.Message{
							{
								CreatedAt: time.Now(),
								Message:   "user_not_found",
							},
						},
					},
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
			Data: []model.UserMessage{
				{
					ID:             user.ID,
					Username:       user.Username,
					Avatar:         &user.Avatar,
					LastActiveDate: time.Now(),
					Messages: []model.Message{
						{
							CreatedAt: time.Now(),
							Message:   "Successfully connected to messaging service",
							Type:      1,
						},
					},
				},
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
			var msg model.WSMessageReceived

			fmt.Println("Timezone check: ", utils.CheckGMTTime(msg.Data.Time))

			if err := c.ReadJSON(&msg); err != nil {
				log.Printf("❌ Read error: %v", err)
				break
			}

			log.Printf("📨 Received message from user %d: %s", user.ID, msg.Event)

			switch msg.Event {
			case "ping":
				pongData := []model.UserMessage{
					{
						Messages: []model.Message{
							{
								CreatedAt:  time.Now(),
								Message:    "pong",
								Type:       1,
								SenderID:   user.ID,
								ReceiverID: user.ID,
								ID:         1,
							},
						},
					},
				}

				pongMsg := model.WSMessage{
					Event: "pong",
					Data:  pongData,
				}
				c.WriteJSON(pongMsg)

			case "private_message":
				targetC, exists := wsUserConns[msg.TargetUserID]

				if !exists || len(targetC) == 0 {
					// todo: write to database that the user is not online
					err := h.service.MessageWriteToDatabase(user.ID, false, msg.Data)

					if err != nil {
						log.Printf("❌ Error writing to database: %v", err)
					}

					log.Printf("❌ User %d not connected", msg.TargetUserID)
					continue
				}

				// targetUser := wsClients[targetC[0]]
				data := []model.UserMessage{
					{
						ID:             user.ID,
						Username:       user.Username,
						Avatar:         &user.Avatar,
						LastActiveDate: time.Now(),
						Messages: []model.Message{
							{
								CreatedAt:  msg.Data.Time,
								Message:    msg.Data.Message,
								Type:       msg.Data.Type,
								SenderID:   user.ID,
								ReceiverID: msg.TargetUserID,
							},
						},
					},
				}

				sendToUser(msg.TargetUserID, "new_message", data)

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

	for i := range connections {
		go func(c *websocket.Conn) {
			if err := c.WriteJSON(msg); err != nil {
				log.Printf("❌ Send error: %v", err)
			}
		}(connections[i])
	}
}
