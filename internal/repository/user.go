package repository

import (
	"backend/internal/domain"

	"github.com/jackc/pgx/v5"
)

type User interface {
	Create(
		tx *pgx.Tx,
		id string,
	) (*domain.User, error)

	Update(
		tx *pgx.Tx,
		userID string,
		firstName *string,
		lastName *string,
		email *string,
		dateOfBirth *string,
		phoneNumber *string,
		address *string,
	) (*domain.User, error)

	List(
		tx *pgx.Tx,
		limit int,
		offset int,
	) (*[]domain.User, error)

	Get(
		tx *pgx.Tx,
		userID string,
	) (*domain.User, error)

	GetCart(
		tx *pgx.Tx,
		userID string,
	) (*Cart, error)
}
