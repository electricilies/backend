package application

import (
	"context"

	"backend/internal/domain/payment"
)

type Payment interface {
	Create(context.Context, *payment.Model) (*payment.Model, error)
	Get(context.Context, int) (*payment.Model, error)
	Update(context.Context, *payment.Model, int) (*payment.Model, error)
	Delete(context.Context, int) error
}

type PaymentImpl struct{}
