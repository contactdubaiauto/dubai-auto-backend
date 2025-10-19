package repository

import (
	"context"
	"dubai-auto/internal/model"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SocketRepository struct {
	db *pgxpool.Pool
}

func NewSocketRepository(db *pgxpool.Pool) *SocketRepository {
	return &SocketRepository{db}
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

func (r *SocketRepository) GetUserAvatar(userID int) (string, error) {
	q := `
		SELECT avatar FROM profiles WHERE user_id = $1
	`
	var avatar string
	err := r.db.QueryRow(context.Background(), q, userID).Scan(&avatar)
	return avatar, err
}

func (r *SocketRepository) MessageWriteToDatabase(senderUserID int, status bool, msg model.MessageReceived) error {
	s := 1

	if status {
		s = 2
	}
	fmt.Println("message data: ", senderUserID, msg.TargetUserID, s, msg.Message, msg.Type, msg.Time)
	q := `
		INSERT INTO messages (
			sender_id, receiver_id, status, message, type, created_at
		) VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.Exec(context.Background(), q, senderUserID, msg.TargetUserID, s, msg.Message, msg.Type, msg.Time)
	return err
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
