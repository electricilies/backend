package service

import (
	"backend/internal/domain"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/hashicorp/go-multierror"
)

type Cart struct {
	validate *validator.Validate
}

func ProvideCart(
	validate *validator.Validate,
) *Cart {
	return &Cart{
		validate: validate,
	}
}

var _ domain.CartService = &Cart{}

func (c *Cart) CreateItem(
	productVariant domain.ProductVariant,
	quantity int,
) (*domain.CartItem, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, multierror.Append(domain.ErrInternal, err)
	}
	cartItem := &domain.CartItem{
		ID:             id,
		ProductVariant: productVariant,
		Quantity:       quantity,
	}
	if err := c.validate.Struct(cartItem); err != nil {
		return nil, multierror.Append(domain.ErrInvalid, err)
	}
	return cartItem, nil
}
