package http

import (
	"dubai-auto/internal/model"
	"dubai-auto/internal/service"
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
	// pending messages awaiting ack; key: "<receiverID>|<RFC3339Nano time>" value: attempts count
	tmpMessages = make(map[string]int)
	tmpMutex    = sync.Mutex{}
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
							Type:      1, // 1-text, 2-item, 3-video, 4-image,
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

		// Heartbeat: emit ping and expect any incoming event within 3s.
		// Retry up to 2 additional times; disconnect if still no response.
		heartbeatCh := make(chan struct{}, 1)
		done := make(chan struct{})

		// Stop heartbeat goroutine when connection handler exits
		defer close(done)

		go func(conn *websocket.Conn) {
			missCount := 0
			for {
				select {
				case <-done:
					return
				default:
				}

				// send ping event to this connection
				_ = conn.WriteJSON(model.WSMessage{Event: "ping"})

				select {
				case <-heartbeatCh:
					// received any activity
					missCount = 0
					// small idle before next heartbeat cycle
					time.Sleep(3 * time.Second)
				case <-time.After(3 * time.Second):
					missCount++
					if missCount >= 3 { // initial try + 2 retries
						log.Printf("⛔ Closing idle websocket for user %d due to heartbeat timeout", user.ID)
						_ = conn.Close()
						return
					}
				}
			}
		}(c)

		for {
			var msg model.WSMessageReceived

			if err := c.ReadJSON(&msg); err != nil {
				log.Printf("❌ Read error: %v", err)
				break
			}

			// any received event counts as a heartbeat response
			select {
			case heartbeatCh <- struct{}{}:
			default:
			}

			switch msg.Event {
			case "ping":
				pongData := []model.UserMessage{
					{
						Messages: []model.Message{
							{
								CreatedAt:  time.Now(),
								Message:    "pong",
								Type:       1, // 1-text, 2-item, 3-video, 4-image,
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
				s := false
				targetC, exists := wsUserConns[msg.TargetUserID]
				messageReceived := msg.Data.(model.MessageReceived)
				sendToUser(user.ID, "ack", messageReceived.Time)

				if exists && len(targetC) > 0 {
					s = true
					data := []model.UserMessage{
						{
							ID:             user.ID,
							Username:       user.Username,
							Avatar:         &user.Avatar,
							LastActiveDate: time.Now(),
							Messages: []model.Message{
								{
									CreatedAt:  messageReceived.Time,
									Message:    messageReceived.Message,
									Type:       messageReceived.Type,
									SenderID:   user.ID,
									ReceiverID: msg.TargetUserID,
								},
							},
						},
					}

					sendToUser(msg.TargetUserID, "new_message", data)

					// store pending message for ack tracking and retry
					key := fmt.Sprintf("%d|%s", msg.TargetUserID, messageReceived.Time.Format(time.RFC3339Nano))
					tmpMutex.Lock()
					tmpMessages[key] = 1 // initial attempt just sent
					tmpMutex.Unlock()

					go func(receiverID int, payload []model.UserMessage, original model.WSMessageReceived) {
						for {
							time.Sleep(3 * time.Second)
							k := fmt.Sprintf("%d|%s", receiverID, messageReceived.Time.Format(time.RFC3339Nano))
							tmpMutex.Lock()
							attempts, exists := tmpMessages[k]

							if !exists {
								// ack received; stop retries
								tmpMutex.Unlock()
								return
							}

							if attempts >= 3 {
								// give up: send push and disconnect receiver, then clear pending
								delete(tmpMessages, k)
								tmpMutex.Unlock()

								if err := h.service.SendPushForMessage(user.ID, messageReceived); err != nil {
									log.Printf("❌ Error sending push: %v", err)
								}

								disconnectUser(receiverID)
								return
							}

							// retry send
							tmpMessages[k] = attempts + 1
							tmpMutex.Unlock()
							sendToUser(receiverID, "new_message", payload)
						}
					}(msg.TargetUserID, data, msg)
				}

				err = h.service.MessageWriteToDatabase(user.ID, s, messageReceived)

				if err != nil {
					log.Printf("❌ Error writing to database: %v", err)
				}

			case "ack":
				// receiver acknowledged message delivery; clear pending attempt tracking
				fmt.Println("ack", msg.Data)
				t := msg.Data.(time.Time)
				fmt.Println("ack - converted to time.Time", t)
				key := fmt.Sprintf("%d|%s", user.ID, t.Format(time.RFC3339Nano))
				tmpMutex.Lock()
				delete(tmpMessages, key)
				tmpMutex.Unlock()

			default:
				log.Printf("⚠️ Unknown event: %s", msg.Event)
			}
		}
	}))
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

func disconnectUser(userID int) {
	wsMutex.Lock()
	conns, exists := wsUserConns[userID]
	if exists {
		for _, conn := range conns {
			_ = conn.Close()
		}
		delete(wsUserConns, userID)
	}
	wsMutex.Unlock()
	log.Printf("🔌 Disconnected user %d due to missing ack", userID)
}
