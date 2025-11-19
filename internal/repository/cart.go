package repository

import (
	"backend/internal/domain"
	"github.com/jackc/pgx/v5"
)

type Cart interface {
	Get(
		tx pgx.Tx,
		id int,
	) (*domain.Cart, error)

	Create(
		tx pgx.Tx,
		userID int,
	) (*domain.Cart, error)

	CreateItem(
		tx pgx.Tx,
		cartID int,
		productVariantID int,
		quantity int,
	) (*domain.CartItem, error)

	UpdateItem(
		tx pgx.Tx,
		cartID int,
		itemID string,
		quantity *int,
	) (*domain.CartItem, error)

	DeleteItem(
		tx pgx.Tx,
		cartID int,
		itemID string,
	) error

	GetItem(
		tx pgx.Tx,
		cartID int,
		itemID string,
	) (*domain.CartItem, error)

	Delete(
		tx pgx.Tx,
		id int,
	) error
}
