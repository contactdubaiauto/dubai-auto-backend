package model

import (
	"time"

	"github.com/gofiber/contrib/websocket"
)

type WSMessage struct {
	Event        string `json:"event"`
	TargetUserID int    `json:"target_user_id"`
	Data         any    `json:"data"`
}

type MessageReceived struct {
	Time         time.Time `json:"time"`
	Message      string    `json:"message"`
	TargetUserID int       `json:"target_user_id"` // todo: remove this field, it is have in WSMessageReceived
	Type         int       `json:"type"`
}

type WSMessageReceived struct {
	Event        string `json:"event"`
	TargetUserID int    `json:"target_user_id"`
	Data         any    `json:"data"`
}

type WSMessageAck struct {
	Event        string    `json:"event"`
	TargetUserID int       `json:"target_user_id"`
	Data         time.Time `json:"data"`
}

type WSUser struct {
	ID       int             `json:"id"`
	Username string          `json:"username"`
	Avatar   string          `json:"avatar"`
	RoleID   int             `json:"role_id"`
	Conn     *websocket.Conn `json:"-"`
}

type Message struct {
	CreatedAt  time.Time `json:"created_at"`
	Message    string    `json:"message"`
	ID         int       `json:"id"`
	SenderID   int       `json:"sender_id"`
	ReceiverID int       `json:"receiver_id"`
	Type       int       `json:"type"`
}

type UserMessage struct {
	Messages       []Message `json:"messages"`
	LastActiveDate time.Time `json:"last_active_date"`
	Username       string    `json:"username"`
	Avatar         *string   `json:"avatar"`
	ID             int       `json:"id"`
}

type FirebaseMessageData struct {
	Time     time.Time `json:"time"`
	Message  string    `json:"message"`
	Username string    `json:"username"`
	Avatar   string    `json:"avatar"`
	Type     int       `json:"type"`
}

type FirebaseMessage struct {
	Data FirebaseMessageData `json:"data"`
	Type int                 `json:"type"`
}
