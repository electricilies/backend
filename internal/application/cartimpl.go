package application

import (
	"context"

	"backend/internal/delivery/http"
	"backend/internal/domain"
)

type Cart struct {
	cartRepo    domain.CartRepository
	cartService domain.CartService
}

func ProvideCart(cartRepo domain.CartRepository, cartService domain.CartService) *Cart {
	return &Cart{
		cartRepo:    cartRepo,
		cartService: cartService,
	}
}

var _ http.CartApplication = &Cart{}

func (c *Cart) Get(ctx context.Context, param http.GetCartRequestDto) (*domain.Cart, error) {
	cart, err := c.cartRepo.Get(ctx, param.CartID)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (c *Cart) Create(ctx context.Context, param http.CreateCartRequestDto) (*domain.Cart, error) {
	cart, err := c.cartService.Create(param.UserID)
	if err != nil {
		return nil, err
	}
	err = c.cartRepo.Save(ctx, *cart)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (c *Cart) CreateItem(ctx context.Context, param http.CreateCartItemRequestDto) (*domain.CartItem, error) {
	cart, err := c.cartRepo.Get(ctx, param.CartID)
	if err != nil {
		return nil, err
	}

	// Verify cart belongs to user
	if cart.UserID != param.UserID {
		return nil, domain.ErrForbidden
	}

	cartItem, err := c.cartService.CreateItem(
		param.Data.ProductID,
		param.Data.ProductVariantID,
		param.Data.Quantity,
	)
	if err != nil {
		return nil, err
	}

	err = c.cartService.AddItem(cart, *cartItem)
	if err != nil {
		return nil, err
	}

	err = c.cartRepo.Save(ctx, *cart)
	if err != nil {
		return nil, err
	}

	return cartItem, nil
}

func (c *Cart) UpdateItem(ctx context.Context, param http.UpdateCartItemRequestDto) (*domain.CartItem, error) {
	cart, err := c.cartRepo.Get(ctx, param.CartID)
	if err != nil {
		return nil, err
	}

	// Verify cart belongs to user
	if cart.UserID != param.UserID {
		return nil, domain.ErrForbidden
	}

	err = c.cartService.UpdateItem(cart, param.ItemID, param.Data.Quantity)
	if err != nil {
		return nil, err
	}

	err = c.cartRepo.Save(ctx, *cart)
	if err != nil {
		return nil, err
	}

	// Find and return updated item
	if cart.Items != nil {
		for _, item := range cart.Items {
			if item.ID == param.ItemID {
				return &item, nil
			}
		}
	}

	return nil, domain.ErrNotFound
}

func (c *Cart) DeleteItem(ctx context.Context, param http.DeleteCartItemRequestDto) error {
	cart, err := c.cartRepo.Get(ctx, param.CartID)
	if err != nil {
		return err
	}

	// Verify cart belongs to user
	if cart.UserID != param.UserID {
		return domain.ErrForbidden
	}

	err = c.cartService.RemoveItem(cart, param.ItemID)
	if err != nil {
		return err
	}

	err = c.cartRepo.Save(ctx, *cart)
	return err
}
