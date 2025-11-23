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

func (c *Cart) Create(
	userID uuid.UUID,
) (*domain.Cart, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, multierror.Append(domain.ErrInternal, err)
	}
	cart := &domain.Cart{
		ID:     id,
		UserID: userID,
	}
	if err := c.validate.Struct(cart); err != nil {
		return nil, multierror.Append(domain.ErrInvalid, err)
	}
	return cart, nil
}

func (c *Cart) CreateItem(
	productVariantID uuid.UUID,
	quantity int,
) (*domain.CartItem, error) {
	// For demonstration, ProductVariant is not loaded from DB
	productVariant := domain.ProductVariant{ID: productVariantID}
	id, err := uuid.NewV7()
	if err != nil {
		return nil, multierror.Append(domain.ErrInternal, err)
	}
	cartItem := &domain.CartItem{
		ID:             id,
		ProductVariant: &productVariant,
		Quantity:       quantity,
	}
	if err := c.validate.Struct(cartItem); err != nil {
		return nil, multierror.Append(domain.ErrInvalid, err)
	}
	return cartItem, nil
}

func (c *Cart) AddItem(
	cart *domain.Cart,
	item domain.CartItem,
) error {
	if cart.Items == nil {
		cart.Items = []domain.CartItem{}
	}
	cart.Items = append(cart.Items, item)
	if err := c.validate.Struct(cart); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}

func (c *Cart) UpdateItem(
	cart *domain.Cart,
	itemID uuid.UUID,
	quantity int,
) error {
	if cart.Items == nil {
		return multierror.Append(domain.ErrInvalid, domain.ErrNotFound)
	}
	for i, item := range cart.Items {
		if item.ID == itemID {
			(cart.Items)[i].Quantity = quantity
			if err := c.validate.Struct(cart); err != nil {
				return multierror.Append(domain.ErrInvalid, err)
			}
			return nil
		}
	}
	return multierror.Append(domain.ErrInvalid, domain.ErrNotFound)
}

func (c *Cart) RemoveItem(
	cart *domain.Cart,
	itemID uuid.UUID,
) error {
	if cart.Items == nil {
		return nil
	}
	newItems := []domain.CartItem{}
	for _, item := range cart.Items {
		if item.ID != itemID {
			newItems = append(newItems, item)
		}
	}
	cart.Items = newItems
	if err := c.validate.Struct(cart); err != nil {
		return multierror.Append(domain.ErrInvalid, err)
	}
	return nil
}
