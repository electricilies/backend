package repository

import (
	"backend/internal/domain"

	"github.com/jackc/pgx/v5"
)

type Payment interface {
	Create(
		tx pgx.Tx,
		provider domain.PaymentProvider,
		orderId int,
	) (*domain.Payment, error)

	Update(
		tx pgx.Tx,
		id int,
		status domain.PaymentStatus,
	) (*domain.Payment, error)

	List(
		tx pgx.Tx,
		limit int,
		offset int,
	) (*[]domain.Payment, error)

	Get(
		tx pgx.Tx,
		id int,
	) (*domain.Payment, error)
}
