package service

import (
	"context"

	"backend/internal/domain"
)

type Payment interface {
	Create(context.Context, CreatePaymentParam) (*domain.Payment, error)
	Update(context.Context, UpdatePaymentParam) (*domain.Payment, error)
	List(context.Context, ListPaymentParam) (*Pagination[domain.Payment], error)
	Get(context.Context, GetPaymentParam) (*domain.Payment, error)
}
