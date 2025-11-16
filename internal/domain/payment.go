package payment

import (
	"time"

	"backend/internal/domain/order"
)

type Provider string

const (
	COD     Provider = "COD"
	VNPAY   Provider = "VNPAY"
	MOMO    Provider = "MOMO"
	ZALOPAY Provider = "ZALOPAY"
)

type Status string

const (
	Pending   Status = "PENDING"
	Completed Status = "COMPLETED"
	Failed    Status = "FAILED"
)

type Model struct {
	ID              int
	Amount          int64
	PaymentStatus   Status
	Order           order.Model
	PaymentProvider Provider
	UpdatedAt       time.Time
}
