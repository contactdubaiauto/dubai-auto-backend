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

	wsUserConns    map[int]*websocket.Conn
	wsMutex        sync.RWMutex
	tmpMessages    map[string]int
	tmpMutex       sync.Mutex
	connWriteMutex map[*websocket.Conn]*sync.Mutex
	connWriteMuMux sync.RWMutex
}

func (h *SocketHandler) closeConnGracefully(c *websocket.Conn) {

	if c == nil {
		return
	}

	h.connWriteMuMux.Lock()
	writeMu, exists := h.connWriteMutex[c]
	if exists && writeMu != nil {
		writeMu.Lock()
		defer writeMu.Unlock()
	}
	h.connWriteMuMux.Unlock()

	_ = c.WriteControl(websocket.CloseMessage, []byte{}, time.Now().Add(1*time.Second))
	_ = c.Close()

	h.connWriteMuMux.Lock()
	delete(h.connWriteMutex, c)
	h.connWriteMuMux.Unlock()
}

func (h *SocketHandler) safeWriteJSON(c *websocket.Conn, msg any) error {
	if c == nil {
		return fmt.Errorf("connection is nil")
	}

	h.connWriteMuMux.RLock()
	writeMu, exists := h.connWriteMutex[c]
	h.connWriteMuMux.RUnlock()

	if !exists {
		h.connWriteMuMux.Lock()
		writeMu, exists = h.connWriteMutex[c]
		if !exists {
			writeMu = &sync.Mutex{}
			h.connWriteMutex[c] = writeMu
		}
		h.connWriteMuMux.Unlock()
	}

	writeMu.Lock()
	defer writeMu.Unlock()

	return c.WriteJSON(msg)
}

type SocketHandlerOption func(*SocketHandler)

func NewSocketHandler(service *service.SocketService, opts ...SocketHandlerOption) *SocketHandler {
	handler := &SocketHandler{
		service:        service,
		wsUserConns:    make(map[int]*websocket.Conn),
		tmpMessages:    make(map[string]int),
		connWriteMutex: make(map[*websocket.Conn]*sync.Mutex),
	}

	for _, opt := range opts {
		if opt != nil {
			opt(handler)
		}
	}

	return handler
}

func (h *SocketHandler) handleHeartbeat(conn *websocket.Conn, userID int, heartbeatTimeout chan struct{}, heartbeatCh chan struct{}, done chan struct{}) {
	missCount := 0
	for {
		select {
		case <-done:
			return
		default:
		}

		_ = h.safeWriteJSON(conn, model.WSMessage{Event: "ping"})

		select {
		case <-heartbeatCh:
			missCount = 0
			time.Sleep(3 * time.Second)
		case <-time.After(3 * time.Second):
			missCount++
			if missCount >= 3 {
				log.Printf("⛔ Closing idle websocket for user %d due to heartbeat timeout", userID)
				heartbeatTimeout <- struct{}{}
				return
			}
		}
	}
}

func (h *SocketHandler) SetupWebSocketHandler() fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
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

			h.safeWriteJSON(c, errMsg)
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
			h.safeWriteJSON(c, errMsg)
			c.Close()
			return
		}

		user.Conn = c
		h.wsMutex.Lock()

		if old, exists := h.wsUserConns[user.ID]; exists && old != nil && old != c {
			h.closeConnGracefully(old)
		}

		h.wsUserConns[user.ID] = c
		h.wsMutex.Unlock()

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

		h.safeWriteJSON(c, welcomeMsg)
		err = h.service.UpdateUserStatus(user.ID, true)

		if err != nil {
			log.Printf("❌ Error updating user status: %v", err)
		}

		defer func() {
			h.wsMutex.Lock()

			if conn, exists := h.wsUserConns[user.ID]; exists {

				if conn == c {
					delete(h.wsUserConns, user.ID)
				}
			}

			h.wsMutex.Unlock()
			err = h.service.UpdateUserStatus(user.ID, false)

			if err != nil {
				log.Printf("❌ Error updating user status: %v", err)
			}

			h.closeConnGracefully(c)
		}()

		data, err := h.service.GetNewMessages(user.ID)

		if err != nil {
			log.Printf("❌ Error getting unread messages: %v", err)
		} else if data != nil {
			h.sendToUser(user.ID, "new_message", data)
		}

		heartbeatCh := make(chan struct{}, 1)
		done := make(chan struct{})
		heartbeatTimeout := make(chan struct{})
		defer close(done)

		go h.handleHeartbeat(c, user.ID, heartbeatTimeout, heartbeatCh, done)

		for {
			var msg model.WSMessageReceived

			select {
			case <-heartbeatTimeout:
				log.Printf("ℹ️ Exiting read loop for user %d after heartbeat timeout", user.ID)
				h.closeConnGracefully(c)
				return
			default:
			}

			if err := c.ReadJSON(&msg); err != nil {
				select {
				case <-heartbeatTimeout:
					log.Printf("ℹ️ Connection closed for user %d due to heartbeat timeout", user.ID)
				default:
					log.Printf("❌ Read error: %v", err)
				}
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

				h.sendToUser(user.ID, "ack", messageReceived.Time)

				if msg.TargetUserID == 0 {
					adminIDs, err := h.service.GetActiveAdminsWithChatPermission()

					if err != nil {
						log.Printf("❌ Error getting admins with chat permission: %v", err)
					} else {
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
										ReceiverID: 0, // Admin broadcast
									},
								},
							},
						}

						// Send to all admins and track delivery
						for _, adminID := range adminIDs {
							h.wsMutex.RLock()
							targetC, exists := h.wsUserConns[adminID]
							h.wsMutex.RUnlock()
							s = false

							if exists && targetC != nil {
								s = true
								h.sendToUser(adminID, "new_message", data)
								key := fmt.Sprintf("%d|%s", adminID, messageReceived.Time.Format(time.RFC3339Nano))
								h.tmpMutex.Lock()
								h.tmpMessages[key] = 1
								h.tmpMutex.Unlock()

								// Retry mechanism for each admin
								go func(receiverID int, payload []model.UserMessage) {
									for {
										time.Sleep(3 * time.Second)
										k := fmt.Sprintf("%d|%s", receiverID, messageReceived.Time.Format(time.RFC3339Nano))
										h.tmpMutex.Lock()
										attempts, exists := h.tmpMessages[k]

										if !exists {
											h.tmpMutex.Unlock()
											return
										}

										if attempts >= 3 {
											delete(h.tmpMessages, k)
											h.tmpMutex.Unlock()
											fmt.Println("attempts for new message:", attempts)

											// Send push notification to this admin
											msgCopy := messageReceived
											msgCopy.TargetUserID = receiverID

											if err := h.service.SendPushForMessage(user.ID, msgCopy); err != nil {
												log.Printf("❌ Error sending push to admin %d: %v", receiverID, err)
											}

											h.closeConnGracefully(targetC)
											return
										}

										h.tmpMessages[k] = attempts + 1
										h.tmpMutex.Unlock()
										h.sendToUser(receiverID, "new_message", payload)
									}
								}(adminID, data)

								msgCopy := messageReceived
								msgCopy.TargetUserID = adminID
								err = h.service.MessageWriteToDatabase(user.ID, s, msgCopy)

								if err != nil {
									log.Printf("❌ Error writing to database for admin %d: %v", adminID, err)
								}
							} else {
								msgCopy := messageReceived
								msgCopy.TargetUserID = adminID
								err = h.service.MessageWriteToDatabase(user.ID, false, msgCopy)

								if err != nil {
									log.Printf("❌ Error writing to database for admin %d: %v", adminID, err)
								}
							}
						}

					}
				} else {
					// Regular private message to specific user
					h.wsMutex.RLock()
					targetC, exists := h.wsUserConns[msg.TargetUserID]
					h.wsMutex.RUnlock()

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

						h.sendToUser(msg.TargetUserID, "new_message", data)
						key := fmt.Sprintf("%d|%s", msg.TargetUserID, messageReceived.Time.Format(time.RFC3339Nano))
						h.tmpMutex.Lock()
						h.tmpMessages[key] = 1
						h.tmpMutex.Unlock()

						go func(receiverID int, payload []model.UserMessage, original model.WSMessageReceived) {
							for {
								time.Sleep(3 * time.Second)
								k := fmt.Sprintf("%d|%s", receiverID, messageReceived.Time.Format(time.RFC3339Nano))
								h.tmpMutex.Lock()
								attempts, exists := h.tmpMessages[k]

								if !exists {
									h.tmpMutex.Unlock()
									return
								}

								if attempts >= 3 {
									delete(h.tmpMessages, k)
									h.tmpMutex.Unlock()
									fmt.Println("attempts for new message:", attempts)

									if err := h.service.SendPushForMessage(user.ID, messageReceived); err != nil {
										log.Printf("❌ Error sending push: %v", err)
									}

									h.closeConnGracefully(targetC)
									return
								}

								h.tmpMessages[k] = attempts + 1
								h.tmpMutex.Unlock()
								h.sendToUser(receiverID, "new_message", payload)
							}
						}(msg.TargetUserID, data, msg)
					}

					err = h.service.MessageWriteToDatabase(user.ID, s, messageReceived)

					if err != nil {
						log.Printf("❌ Error writing to database: %v", err)
					}
				}

			case "ack":

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
				h.tmpMutex.Lock()
				delete(h.tmpMessages, key)
				h.tmpMutex.Unlock()

			default:
				log.Printf("⚠️ Unknown event: %s", msg.Event)
			}
		}
		fmt.Println("for is ended: ", user.ID)
	})
}

func (h *SocketHandler) sendToUser(userID int, event string, data any) {
	h.wsMutex.RLock()
	userConn, exists := h.wsUserConns[userID]

	if !exists || userConn == nil {
		h.wsMutex.RUnlock()
		log.Printf("❌ User %d not connected", userID)
		return
	}

	h.wsMutex.RUnlock()
	log.Printf("📤 Sending %s to user %d", event, userID)

	msg := model.WSMessage{
		Event:        event,
		TargetUserID: userID,
		Data:         data,
	}

	go h.safeWriteJSON(userConn, msg)
}
