package http

import (
	"dubai-auto/internal/model"
	"dubai-auto/internal/service"
	"dubai-auto/pkg/auth"
	"encoding/json"
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

// todo: update the handler file :(
func closeConnGracefully(c *websocket.Conn) {
	if c == nil {
		return
	}

	// Get write mutex before closing
	connWriteMuMux.Lock()
	writeMu, exists := connWriteMutex[c]
	if exists && writeMu != nil {
		writeMu.Lock()
		defer writeMu.Unlock()
	}
	connWriteMuMux.Unlock()

	_ = c.WriteControl(websocket.CloseMessage, []byte{}, time.Now().Add(1*time.Second))
	_ = c.Close()

	// Clean up write mutex
	connWriteMuMux.Lock()
	delete(connWriteMutex, c)
	connWriteMuMux.Unlock()
}

// safeWriteJSON safely writes JSON to a WebSocket connection using a per-connection mutex
func safeWriteJSON(c *websocket.Conn, msg interface{}) error {
	if c == nil {
		return fmt.Errorf("connection is nil")
	}

	connWriteMuMux.RLock()
	writeMu, exists := connWriteMutex[c]
	connWriteMuMux.RUnlock()

	if !exists {
		// Create write mutex for this connection if it doesn't exist
		connWriteMuMux.Lock()
		writeMu, exists = connWriteMutex[c]
		if !exists {
			writeMu = &sync.Mutex{}
			connWriteMutex[c] = writeMu
		}
		connWriteMuMux.Unlock()
	}

	writeMu.Lock()
	defer writeMu.Unlock()

	return c.WriteJSON(msg)
}

func NewSocketHandler(service *service.SocketService) *SocketHandler {
	return &SocketHandler{service}
}

var (
	wsClients      = make(map[*websocket.Conn]*model.WSUser)
	wsUserConns    = make(map[int]*websocket.Conn)
	wsMutex        = sync.RWMutex{}
	tmpMessages    = make(map[string]int)
	tmpMutex       = sync.Mutex{}
	connWriteMutex = make(map[*websocket.Conn]*sync.Mutex)
	connWriteMuMux = sync.RWMutex{}
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

			safeWriteJSON(c, errMsg)
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
			safeWriteJSON(c, errMsg)
			c.Close()
			return
		}

		user.Conn = c
		wsMutex.Lock()

		if old, exists := wsUserConns[user.ID]; exists && old != nil && old != c {
			closeConnGracefully(old)
		}

		wsUserConns[user.ID] = c
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
							Type:      1,
						},
					},
				},
			},
		}

		avatar, username := h.service.GetUserAvatar(user.ID)
		user.Avatar = avatar
		user.Username = username
		wsClients[c] = user

		safeWriteJSON(c, welcomeMsg)
		err = h.service.UpdateUserStatus(user.ID, true)

		if err != nil {
			log.Printf("❌ Error updating user status: %v", err)
		}

		defer func() {
			wsMutex.Lock()

			if conn, exists := wsUserConns[user.ID]; exists {

				if conn == c {
					delete(wsUserConns, user.ID)
				}
			}

			wsMutex.Unlock()
			err = h.service.UpdateUserStatus(user.ID, false)

			if err != nil {
				log.Printf("❌ Error updating user status: %v", err)
			}

			closeConnGracefully(c)
		}()

		data, err := h.service.GetNewMessages(user.ID)

		if err != nil {
			log.Printf("❌ Error getting unread messages: %v", err)
		} else if data != nil {
			sendToUser(user.ID, "new_message", data)
		}

		heartbeatCh := make(chan struct{}, 1)
		done := make(chan struct{})
		defer close(done)

		go func(conn *websocket.Conn) {
			missCount := 0
			for {
				select {
				case <-done:
					return
				default:
				}

				_ = safeWriteJSON(conn, model.WSMessage{Event: "ping"})

				select {
				case <-heartbeatCh:
					missCount = 0
					time.Sleep(3 * time.Second)
				case <-time.After(3 * time.Second):
					missCount++
					fmt.Println("missCoun")
					fmt.Println(missCount)
					if missCount >= 3 {
						log.Printf("⛔ Closing idle websocket for user %d due to heartbeat timeout", user.ID)
						closeConnGracefully(conn)
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

			select {
			case heartbeatCh <- struct{}{}:
			default:
			}

			switch msg.Event {
			case "ping":

			case "private_message":
				s := false
				targetC, exists := wsUserConns[msg.TargetUserID]
				var messageReceived model.MessageReceived

				if b, err := json.Marshal(msg.Data); err == nil {

					if err := json.Unmarshal(b, &messageReceived); err != nil {
						log.Printf("❌ Error decoding private_message data: %v", err)
						break
					}

				} else {
					log.Printf("❌ Error marshaling private_message data: %v", err)
					break
				}

				sendToUser(user.ID, "ack", messageReceived.Time)

				if exists && targetC != nil {
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
					key := fmt.Sprintf("%d|%s", msg.TargetUserID, messageReceived.Time.Format(time.RFC3339Nano))
					tmpMutex.Lock()
					tmpMessages[key] = 1
					tmpMutex.Unlock()

					go func(receiverID int, payload []model.UserMessage, original model.WSMessageReceived) {
						for {
							time.Sleep(3 * time.Second)
							k := fmt.Sprintf("%d|%s", receiverID, messageReceived.Time.Format(time.RFC3339Nano))
							tmpMutex.Lock()
							attempts, exists := tmpMessages[k]

							if !exists {
								tmpMutex.Unlock()
								return
							}

							if attempts >= 3 {
								delete(tmpMessages, k)
								tmpMutex.Unlock()
								fmt.Println("attempts for new message:", attempts)

								if err := h.service.SendPushForMessage(user.ID, messageReceived); err != nil {
									log.Printf("❌ Error sending push: %v", err)
								}

								disconnectUser(receiverID)
								return
							}

							tmpMessages[k] = attempts + 1
							tmpMutex.Unlock()
							sendToUser(receiverID, "new_message", payload)
						}
					}(msg.TargetUserID, data, msg)
				}

				fmt.Println("new message time", messageReceived.Time)

				err = h.service.MessageWriteToDatabase(user.ID, s, messageReceived)

				if err != nil {
					log.Printf("❌ Error writing to database: %v", err)
				}

			case "ack":

				fmt.Println("ack event", msg.Data)
				var t time.Time
				switch v := msg.Data.(type) {
				case string:
					parsed, err := time.Parse(time.RFC3339Nano, v)
					if err != nil {

						parsed, err = time.Parse(time.RFC3339, v)
					}
					if err != nil {
						log.Printf("❌ Error parsing ack time: %v", err)
						break
					}
					t = parsed
				case map[string]any:

					if s, ok := v["time"].(string); ok {
						parsed, err := time.Parse(time.RFC3339Nano, s)
						if err != nil {
							parsed, err = time.Parse(time.RFC3339, s)
						}
						if err != nil {
							log.Printf("❌ Error parsing ack time from map: %v", err)
							break
						}
						t = parsed
					} else {
						log.Printf("❌ Unexpected ack data shape: %#v", v)
						break
					}
				case float64:

					t = time.Unix(int64(v), 0)
				default:
					log.Printf("❌ Unsupported ack data type: %T", v)
				}
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
	userConn, exists := wsUserConns[userID]

	if !exists || userConn == nil {
		wsMutex.RUnlock()
		log.Printf("❌ User %d not connected", userID)
		return
	}

	wsMutex.RUnlock()
	log.Printf("📤 Sending %s to user %d", event, userID)

	msg := model.WSMessage{
		Event:        event,
		TargetUserID: userID,
		Data:         data,
	}

	go func(c *websocket.Conn) {

		if err := safeWriteJSON(c, msg); err != nil {
			log.Printf("❌ Send error: %v", err)
		}
	}(userConn)
}

func disconnectUser(userID int) {
	wsMutex.Lock()
	conn, exists := wsUserConns[userID]

	if exists {
		fmt.Println("user disconnectiong ")
		if conn != nil {
			closeConnGracefully(conn)
		}
		delete(wsUserConns, userID)
	}
	wsMutex.Unlock()
}
