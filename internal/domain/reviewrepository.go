package domain

import (
	"github.com/jackc/pgx/v5"
)

type ReviewRepository interface {
	Get(
		tx pgx.Tx,
		id int,
	) (*Review, error)

	List(
		tx pgx.Tx,
		orderItemIDs *[]int,
		productVariantID *int,
		userIDs *[]int,
		deleted string,
		limit int,
		offset int,
	) (*[]Review, error)

	Count(
		tx pgx.Tx,
		orderItemIDs *[]int,
		productVariantID *int,
		userIDs *[]int,
		deleted string,
		limit int,
		offset int,
	) (*int, error)

	Create(
		tx pgx.Tx,
		orderItemID int,
		userID int,
		rating int,
		content string,
		imageURL string,
	) (*Review, error)

	Update(
		tx pgx.Tx,
		id int,
		rating *int,
		content *string,
		imageURL *string,
	) (*Review, error)

	Delete(
		tx pgx.Tx,
		id int,
	) error
}
