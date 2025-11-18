package domain

import (
	"github.com/jackc/pgx/v5"
)

type CategoryRepository interface {
	List(
		tx pgx.Tx,
		search *string,
		limit int,
		offset int,
	) ([]Category, error)
	Get(
		tx pgx.Tx,
		id int,
	) (Category, error)
	Create(
		Tx pgx.Tx,
		Name string,
	) (Category, error)
	Update(
		Tx pgx.Tx,
		id int,
		name *string,
	) (Category, error)
}
