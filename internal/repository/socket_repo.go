package repository

import (
	"context"
	"dubai-auto/internal/config"
	"dubai-auto/internal/model"
	"dubai-auto/pkg/firebase"
	"fmt"

	"firebase.google.com/go/v4/messaging"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SocketRepository struct {
	db              *pgxpool.Pool
	firebaseService *firebase.FirebaseService
	config          *config.Config
}

func NewSocketRepository(db *pgxpool.Pool, firebaseService *firebase.FirebaseService, config *config.Config) *SocketRepository {
	return &SocketRepository{db, firebaseService, config}
}

func (r *SocketRepository) UpdateUserStatus(userID int, status bool) error {
	q := `
		UPDATE users 
		SET online = $1, last_active_date = now() 
		WHERE id = $2
	`
	_, err := r.db.Exec(context.Background(), q, status, userID)
	return err
}

func (r *SocketRepository) GetNewMessages(userID int) ([]model.UserMessage, error) {
	q := `
		WITH updated_messages AS (
			UPDATE messages
			SET status = 2
			WHERE status = 1 AND receiver_id = $1
			RETURNING id, sender_id, receiver_id, message, type, created_at
		)
		SELECT 
			u.id,
			u.username,
			u.last_active_date,
			p.avatar,
			json_agg(
				json_build_object(
					'id', m.id,
					'message', m.message,
					'type', m.type,
					'created_at', to_char(m.created_at AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS"Z"'),
					'sender_id', m.sender_id,
					'receiver_id', m.receiver_id
				)
			) as messages
		FROM updated_messages m
		LEFT JOIN users u ON m.sender_id = u.id
		LEFT JOIN profiles p ON u.id = p.user_id
		GROUP BY u.id, p.avatar;
	`
	rows, err := r.db.Query(context.Background(), q, userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var messages []model.UserMessage

	for rows.Next() {
		var message model.UserMessage
		err := rows.Scan(&message.ID, &message.Username, &message.LastActiveDate, &message.Avatar, &message.Messages)

		if err != nil {
			return messages, err
		}
		messages = append(messages, message)
	}

	return messages, err
}

func (r *SocketRepository) GetUserAvatar(userID int) (string, string, error) {
	q := `
		SELECT avatar, username FROM profiles WHERE user_id = $1
	`
	var avatar, username string
	err := r.db.QueryRow(context.Background(), q, userID).Scan(&avatar, &username)
	return avatar, username, err
}

func (r *SocketRepository) MessageWriteToDatabase(senderUserID int, status bool, msg model.MessageReceived) error {
	s := 1

	if status {
		s = 2
	}

	q := `
		INSERT INTO messages (
			sender_id, receiver_id, status, message, type, created_at
		) VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.Exec(context.Background(), q, senderUserID, msg.TargetUserID, s, msg.Message, msg.Type, msg.Time)
	if err != nil {
		fmt.Println("msg.TargetUserID")
		fmt.Println(msg.TargetUserID)
		return err
	}

	userFcmToken := ""
	q = `
		select device_token from user_tokens where user_id = $1
	`
	r.db.QueryRow(context.Background(), q, msg.TargetUserID).Scan(&userFcmToken)
	username, avatar, _ := r.GetUserAvatarName(senderUserID)

	if !status {
		_, err = r.firebaseService.SendToToken(userFcmToken, messaging.Notification{
			Title:    username,
			Body:     msg.Message,
			ImageURL: avatar,
		})

		if err != nil {
			fmt.Println("error sending notification: ", err)
		}
	}

	return nil
}

func (r *SocketRepository) GetUserAvatarName(userID int) (string, string, error) {
	q := `
		SELECT 
			username,
			$2 || avatar
		FROM profile 
		WHERE user_id = $1
	`
	var username string
	var avatar string
	err := r.db.QueryRow(context.Background(), q, userID, r.config.IMAGE_BASE_URL).Scan(&username, &avatar)
	return username, avatar, err
}

func (r *SocketRepository) CheckUserExists(userID int) error {
	q := `
		SELECT id FROM users WHERE id = $1
	`
	var id int
	err := r.db.QueryRow(context.Background(), q, userID).Scan(&id)
	fmt.Println(id)
	return err
}

func (r *SocketRepository) GetUserToken(userID int) (string, error) {
	q := `
		SELECT device_token FROM user_tokens WHERE user_id = $1
	`
	var token string
	err := r.db.QueryRow(context.Background(), q, userID).Scan(&token)
	return token, err
}

// SendPushForMessage sends a push notification for a given message without writing it to DB.
func (r *SocketRepository) SendPushForMessage(senderUserID int, msg model.MessageReceived) error {
	token, err := r.GetUserToken(msg.TargetUserID)
	if err != nil {
		return err
	}
	username, avatar, _ := r.GetUserAvatarName(senderUserID)
	_, err = r.firebaseService.SendToToken(token, messaging.Notification{
		Title:    username,
		Body:     msg.Message,
		ImageURL: avatar,
	})
	return err
}

// GetActiveAdminsWithChatPermission returns IDs of active admin users with "chat" permission
func (r *SocketRepository) GetActiveAdminsWithChatPermission() ([]int, error) {
	q := `
		SELECT id 
		FROM users 
		WHERE role_id = 0 
		AND status = 1 
		AND permissions @> '["chat"]'::jsonb
	`
	rows, err := r.db.Query(context.Background(), q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var adminIDs []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		adminIDs = append(adminIDs, id)
	}

	return adminIDs, nil
}
