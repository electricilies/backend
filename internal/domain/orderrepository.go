package domain

import (
	"github.com/jackc/pgx/v5"
)

type OrderRepository interface {
	Get(
		tx pgx.Tx,
		id int,
	) (*Order, error)

	List(
		tx pgx.Tx,
		ids *[]int,
		search *string,
		deleted string,
		limit int,
		offset int,
	) (*[]Order, error)

	Create(
		tx pgx.Tx,
		userID int,
		address string,
		productVariantIDs []int,
		quantities []int,
		prices []int64,
	) (*Order, error)

	Update(
		tx pgx.Tx,
		id int,
		status *OrderStatus,
		address *string,
	) (*Order, error)

	Delete(
		tx pgx.Tx,
		id int,
	) error
}
