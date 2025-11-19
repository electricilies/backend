package repository

import (
	"backend/internal/domain"

	"github.com/jackc/pgx/v5"
)

type Review interface {
	Get(
		tx pgx.Tx,
		id int,
	) (*domain.Review, error)

	List(
		tx pgx.Tx,
		orderItemIDs *[]int,
		productVariantID *int,
		userIDs *[]int,
		deleted string,
		limit int,
		offset int,
	) (*[]domain.Review, error)

	Create(
		tx pgx.Tx,
		orderItemID int,
		userID int,
		rating int,
		content string,
		imageURL string,
	) (*domain.Review, error)

	Update(
		tx pgx.Tx,
		id int,
		rating *int,
		content *string,
		imageURL *string,
	) (*domain.Review, error)

	Delete(
		tx pgx.Tx,
		id int,
	) error
}
