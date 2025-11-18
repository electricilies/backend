package domain

import "github.com/jackc/pgx/v5"

type PaymentRepository interface {
	Create(tx pgx.Tx, provider PaymentProvider, orderId int) (*Payment, error)
	Update(tx pgx.Tx, id int, status PaymentStatus) (*Payment, error)
	List(tx pgx.Tx, limit int, offset int) (*[]Payment, error)
	Get(tx pgx.Tx, id int) (*Payment, error)
}
