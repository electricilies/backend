package domain

import (
	"context"
)

type PaymentService interface {
	Create(context.Context, CreatePaymentParam) (*Payment, error)
	Update(context.Context, UpdatePaymentParam) (*Payment, error)
	Get(context.Context, GetPaymentParam) (*Payment, error)
}
