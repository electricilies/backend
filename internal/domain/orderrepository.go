package domain

import (
	"github.com/jackc/pgx/v5"
)

type OrderRepository interface {
	Count(
		tx pgx.Tx,
		ids *[]int,
		deleted string,
		limit int,
		offset int,
	) (*int, error)

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
		provider OrderProvider,
		totalAmount int64,
	) (*Order, error)

	CreateItems(
		tx pgx.Tx,
		orderID int,
		items []struct {
			productVariantIDs int
			quantities        int
			prices            int64
		},
	) (*[]OrderItem, error)

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
