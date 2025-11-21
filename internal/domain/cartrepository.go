package domain

import (
	"github.com/jackc/pgx/v5"
)

type CartRepository interface {
	Get(
		tx pgx.Tx,
		id int,
	) (*Cart, error)

	Create(
		tx pgx.Tx,
		userID int,
	) (*Cart, error)

	CreateItem(
		tx pgx.Tx,
		cartID int,
		productVariantID int,
		quantity int,
	) (*CartItem, error)

	UpdateItem(
		tx pgx.Tx,
		cartID int,
		itemID string,
		quantity *int,
	) (*CartItem, error)

	DeleteItem(
		tx pgx.Tx,
		cartID int,
		itemID string,
	) error

	GetItem(
		tx pgx.Tx,
		cartID int,
		itemID string,
	) (*CartItem, error)
}
