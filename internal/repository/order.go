package repository

import (
	"backend/internal/domain"

	"github.com/jackc/pgx/v5"
)

type Order interface {
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
	) (*domain.Order, error)

	List(
		tx pgx.Tx,
		ids *[]int,
		search *string,
		deleted string,
		limit int,
		offset int,
	) (*[]domain.Order, error)

	Create(
		tx pgx.Tx,
		userID int,
		address string,
		provider domain.OrderProvider,
		totalAmount int64,
	) (*domain.Order, error)

	CreateItems(
		tx pgx.Tx,
		orderID int,
		items []struct {
			productVariantIDs int
			quantities        int
			prices            int64
		},
	) (*[]domain.OrderItem, error)

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
