package model

import (
	"time"

	"github.com/gofiber/contrib/websocket"
)

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
	ID             int       `json:"id"`
}
