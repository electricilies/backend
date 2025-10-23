package db

import (
	"backend/config"
	"backend/internal/infrastructure/presistence/postgres"
	"context"
	"fmt"
	"log"

	"github.com/Thiht/transactor"
	transactorpgx "github.com/Thiht/transactor/pgx"
	"github.com/jackc/pgx/v5"
)

func NewConnection() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		config.Cfg.DBUsername,
		config.Cfg.DBPassword,
		config.Cfg.DBHost,
		config.Cfg.DBPort,
		config.Cfg.DBName,
	))
	if err != nil {
		log.Printf("Cannot connect to DB: %v", err)
		return nil
	}
	return conn
}

func New(c *pgx.Conn) *postgres.Queries {
	q := postgres.New(c)
	return q
}

func NewTransactor(c *pgx.Conn) transactor.Transactor {
	t, _ := transactorpgx.NewTransactor(c)
	return t
}
