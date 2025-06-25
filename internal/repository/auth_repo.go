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

func (r *AuthRepository) UserLogin(ctx context.Context, user *model.UserLoginRequest) (*model.UserByEmail, error) {
	query := `
		SELECT id, email, password FROM users WHERE email = $1
	`
	row := r.db.QueryRow(ctx, query, user.Email)

	var userByEmail model.UserByEmail
	err := row.Scan(&userByEmail.ID, &userByEmail.Email, &userByEmail.Password)

	return &userByEmail, err
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
