package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	"empty/internal/config"
)

func Init() *pgxpool.Pool {
	connectionString := buildConnectionString()
	config, err := pgxpool.ParseConfig(connectionString)

	if err != nil {
		log.Fatalf("Unable to parse database config💊: %v\n", err)
	}

	config.MaxConns = 200
	pool, err := pgxpool.NewWithConfig(context.Background(), config)

	if err != nil {
		log.Fatalf("failed to create connection poolpool🏊: %v\n", err)
	}

	if err = pool.Ping(context.Background()); err != nil {
		panic(fmt.Sprintf("Could not ping postgres🫙 database: %v", err))
	}

	log.Println("Database 🥳 connection pool initialized successfully ✅")
	return pool
}

func buildConnectionString() string {
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		config.ENV.DB_USER, config.ENV.DB_PASSWORD,
		config.ENV.DB_HOST, config.ENV.DB_PORT, config.ENV.DB_NAME,
	)
}
