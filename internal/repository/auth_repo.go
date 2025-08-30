package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valyala/fasthttp"

	"dubai-auto/internal/model"
)

type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{db}
}

func (r *AuthRepository) DeleteAccount(ctx *fasthttp.RequestCtx, userID int) error {

	// Delete user
	_, err := r.db.Exec(ctx, `DELETE FROM users WHERE id = $1`, userID)
	return err
}

func (r *AuthRepository) UserByEmail(ctx *fasthttp.RequestCtx, email *string) (*model.UserByEmail, error) {

	query := `
		SELECT id, email, password, username FROM temp_users WHERE email = $1
	`
	row := r.db.QueryRow(ctx, query, email)

	var u model.UserByEmail
	err := row.Scan(&u.ID, &u.Email, &u.OTP, &u.Username)

	return &u, err
}

func (r *AuthRepository) UserByPhone(ctx *fasthttp.RequestCtx, phone *string) (model.UserByPhone, error) {
	query := `
		SELECT id, phone, password, username FROM temp_users WHERE phone = $1
	`
	row := r.db.QueryRow(ctx, query, phone)

	var userByPhone model.UserByPhone
	err := row.Scan(&userByPhone.ID, &userByPhone.Phone, &userByPhone.OTP, &userByPhone.Username)

	return userByPhone, err
}

func (r *AuthRepository) TempUserEmailGetOrRegister(ctx *fasthttp.RequestCtx, username, email, password string) error {
	var userID int
	q := `
		insert into temp_users (email, password, username, registered_by)
		values ($1, $2, $3, 'email')
		on conflict (email)
		do update
		set 
			password = EXCLUDED.password
		returning id
	`
	err := r.db.QueryRow(ctx, q, email, password, username).Scan(&userID)

	return err
}

func (r *AuthRepository) TempUserPhoneGetOrRegister(ctx *fasthttp.RequestCtx, username, phone, password string) error {
	var userID int
	q := `
		insert into temp_users (phone, password, username, registered_by)
		values ($1, $2, $3, 'phone')
		on conflict (phone)
		do update
		set 
			password = EXCLUDED.password
		returning id
	`
	err := r.db.QueryRow(ctx, q, phone, password, username).Scan(&userID)

	return err
}

func (r *AuthRepository) UserEmailGetOrRegister(ctx *fasthttp.RequestCtx, username, email, password string) (int, error) {
	var userID int
	q := `
		insert into users (email, password)
		values ($1, $2)
		on conflict (email)
		do update
		set 
			password = EXCLUDED.password
		returning id
	`
	err := r.db.QueryRow(ctx, q, email, password).Scan(&userID)

	if err != nil {
		return userID, err
	}

	q = `
		INSERT INTO profiles (user_id, username, registered_by)
		VALUES ($1, $2, $3)
		ON CONFLICT (user_id) DO NOTHING;
	`
	_, err = r.db.Exec(ctx, q, userID, username, "email")
	return userID, err
}

func (r *AuthRepository) UserPhoneGetOrRegister(ctx *fasthttp.RequestCtx, username, phone, password string) (int, error) {

	var userID int
	q := `
		insert into users (phone, password)
		values ($1, $2)
		on conflict (phone)
		do update
		set 
			password = EXCLUDED.password
		returning id
	`
	err := r.db.QueryRow(ctx, q, phone, password).Scan(&userID)

	if err != nil {
		return userID, err
	}

	q = `
		INSERT INTO profiles (user_id, username, registered_by)
		VALUES ($1, $2, $3)
		ON CONFLICT (user_id) DO NOTHING;
	`
	_, err = r.db.Exec(ctx, q, userID, username, "phone")

	return userID, err
}
