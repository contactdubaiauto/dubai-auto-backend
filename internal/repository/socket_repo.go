package repository

import (
	"context"
	"dubai-auto/internal/model"

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
		select 
			u.id,
			u.username,
			u.last_active_date,
			p.avatar,
			json_agg(
				json_build_object(
					'id', m.id,
					'message', m.message,
					'type', m.type,
					'created_at', m.created_at,
					'sender_id', m.sender_id,
					'receiver_id', m.receiver_id
				)
			) as messages
		from messages m
		left join users u on m.sender_id = u.id
		left join profiles p on u.id = p.user_id
		where m.status = 1 and m.receiver_id = $1
		group by u.id, p.avatar;
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
