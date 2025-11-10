package db

import (
	"context"
	"fmt"
	"log"

	"backend/config"
	"backend/internal/infrastructure/presistence/postgres"

	"github.com/Thiht/transactor"
	transactorpgx "github.com/Thiht/transactor/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewConnection(cfg *config.Config) *pgxpool.Pool {
	conn, err := pgxpool.New(context.Background(), fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.DbUsername,
		cfg.DbPassword,
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbName,
	))
	if err != nil {
		log.Printf("Cannot connect to Db: %v", err)
		return nil
	}
	return conn
}

func New(c *pgxpool.Pool) *postgres.Queries {
	q := postgres.New(c)
	return q
}

func NewTransactor(c *pgxpool.Pool) transactor.Transactor {
	t, _ := transactorpgx.NewTransactorFromPool(c)
	return t
}
