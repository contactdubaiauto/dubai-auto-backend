package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type SocketRepository struct {
	db *pgxpool.Pool
}

func NewSocketRepository(db *pgxpool.Pool) *SocketRepository {
	return &SocketRepository{db}
}
