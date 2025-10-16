package model

import (
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
