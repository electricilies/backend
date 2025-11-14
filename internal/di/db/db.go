package db

import (
	"context"
	"log"

	"backend/config"
	"backend/internal/infrastructure/persistence/postgres"

	"github.com/Thiht/transactor"
	transactorpgx "github.com/Thiht/transactor/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewConnection(ctx context.Context, cfg *config.Config) *pgxpool.Pool {
	conn, err := pgxpool.New(ctx, cfg.DBURL)
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
