package application

import (
	"context"

	"backend/internal/domain"
)

type Payment interface {
	List(context.Context, ListPaymentParam) (*Pagination[domain.Payment], error)
}
