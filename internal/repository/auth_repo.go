package repository

import (
	"context"
	"dubai-auto/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{db}
}

func (r *AuthRepository) UserByMail(ctx context.Context, email *string) (*model.UserByEmail, error) {
	query := `
		SELECT id, email, password FROM users WHERE email = $1
	`
	row := r.db.QueryRow(ctx, query, email)

	var userByEmail model.UserByEmail
	err := row.Scan(&userByEmail.ID, &userByEmail.Email, &userByEmail.OTP)

	return &userByEmail, err
}

func (r *AuthRepository) UserByPhone(ctx context.Context, phone *string) (model.UserByPhone, error) {
	query := `
		SELECT id, phone, password FROM users WHERE phone = $1
	`
	row := r.db.QueryRow(ctx, query, phone)

	var userByPhone model.UserByPhone
	err := row.Scan(&userByPhone.ID, &userByPhone.Phone, &userByPhone.OTP)

	return userByPhone, err
}

func (r *AuthRepository) UserMailGetOrRegister(ctx context.Context, username, email, password string) error {
	q := `
		insert into users (username, email, password)
		values ($1, $2, $3)
		on conflict (email)
		do update
		set 
			password = EXCLUDED.password
	`
	_, err := r.db.Exec(ctx, q, username, email, password)

	return err
}

func (r *AuthRepository) UserRegister(ctx context.Context, user *model.UserRegisterRequest) (*int, error) {
	query := `
		INSERT INTO users (email, password, phone, username)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	var id int

	err := r.db.QueryRow(ctx, query, user.Email, user.Password, user.Phone, user.Username).Scan(&id)

	return &id, err
}
