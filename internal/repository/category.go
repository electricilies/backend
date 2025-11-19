package repository

import (
	"backend/internal/domain"
	"github.com/jackc/pgx/v5"
)

type Category interface {
	List(
		tx pgx.Tx,
		search *string,
		limit int,
		offset int,
	) (*[]domain.Category, error)

	Get(
		tx pgx.Tx,
		id int,
	) (*domain.Category, error)

	Create(
		tx pgx.Tx,
		name string,
	) (*domain.Category, error)

	Update(
		tx pgx.Tx,
		id int,
		name *string,
	) (*domain.Category, error)

	Delete(
		tx pgx.Tx,
		id int,
	) error
}
