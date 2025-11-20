package repository

import (
	"github.com/jackc/pgx/v5"
)

type Product interface {
	Create(
		tx pgx.Tx,
		name string,
		description string,
		attributeValueIDs []int,
		categoryID int,
	) (int, error)

	CreateOptions(
		tx pgx.Tx,
		productID int,
		options []struct {
			name   string
			values []string
		},
	) (int, error)

	CreateImages()
}
