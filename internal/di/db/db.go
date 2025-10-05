package db

import (
	"backend/config"
	"backend/internal/infrastructure/presistence/postgres"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func NewDBConnection() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		config.Cfg.DBUsername,
		config.Cfg.DBPassword,
		config.Cfg.DBHost,
		config.Cfg.DBPort,
		config.Cfg.DBName,
	))
	if err != nil {
		log.Fatalf("Cannot connect to DB: %v", err)
	}
	return conn
}

func NewDB(c *pgx.Conn) *postgres.Queries {
	q := postgres.New(c)
	return q
}
