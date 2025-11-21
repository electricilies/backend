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
	) (*[]Category, error)

	Count(
		tx pgx.Tx,
		limit int,
		offset int,
	) (*int, error)

	Get(
		tx pgx.Tx,
		id int,
	) (*Category, error)

	Create(
		tx pgx.Tx,
		name string,
	) (*Category, error)

	Update(
		tx pgx.Tx,
		id int,
		name *string,
	) (*Category, error)

	Delete(
		tx pgx.Tx,
		id int,
	) error
}
