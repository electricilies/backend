package repository

import (
	"backend/internal/domain"
	"github.com/jackc/pgx/v5"
)

type Order interface {
	Get(
		tx pgx.Tx,
		id int,
	) (*domain.Order, error)

	List(
		tx pgx.Tx,
		ids *[]int,
		search *string,
		deleted DeletedParam,
		limit int,
		offset int,
	) (*[]domain.Order, error)

	Create(
		tx pgx.Tx,
		userID int,
		address string,
		productVariantIDs []int,
		quantities []int,
		prices []int64,
	) (*domain.Order, error)

	Update(
		tx pgx.Tx,
		id int,
		status *domain.OrderStatus,
		address *string,
	) (*domain.Order, error)

	Delete(
		tx pgx.Tx,
		id int,
	) error
}
