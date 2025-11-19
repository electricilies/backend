package domain

import "github.com/jackc/pgx/v5"

type UserRepository interface {
	Create(
		tx *pgx.Tx,
		id string,
	) (*User, error)

	Update(
		tx *pgx.Tx,
		userID string,
		firstName *string,
		lastName *string,
		email *string,
		dateOfBirth *string,
		phoneNumber *string,
		address *string,
	) (*User, error)

	List(
		tx *pgx.Tx,
		limit int,
		offset int,
	) (*[]User, error)

	Get(
		tx *pgx.Tx,
		userID string,
	) (*User, error)

	GetCart(
		tx *pgx.Tx,
		userID string,
	) (*Cart, error)
}
