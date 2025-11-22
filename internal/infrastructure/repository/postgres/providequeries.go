package postgres

import "github.com/jackc/pgx/v5/pgxpool"

func ProvideQueries(c *pgxpool.Pool) *Queries {
	q := New(c)
	return q
}
