package client

import (
	"context"
	"log"

	"backend/config"
	"backend/internal/infrastructure/persistence/postgres"

	"github.com/Thiht/transactor"
	transactorpgx "github.com/Thiht/transactor/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDBConnection(ctx context.Context, srvCfg *config.Server) *pgxpool.Pool {
	conn, err := pgxpool.New(ctx, srvCfg.DBURL)
	if err != nil {
		log.Printf("Cannot connect to Db: %v", err)
		return nil
	}
	return conn
}

func NewDBQueries(c *pgxpool.Pool) *postgres.Queries {
	q := postgres.New(c)
	return q
}

func NewDBTransactor(c *pgxpool.Pool) transactor.Transactor {
	t, _ := transactorpgx.NewTransactorFromPool(c)
	return t
}
