package db

import (
	"backend/config"
	"backend/internal/infrastructure/presistence/postgres"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func NewDB() *postgres.Queries {
	conn, err := pgx.Connect(context.Background(), config.Cfg.DBUrl)
	if err != nil {
		log.Fatalf("")
	}
	q := postgres.New(conn)
	return q
}
