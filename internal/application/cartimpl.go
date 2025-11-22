package application

import (
	"context"

	"backend/internal/domain"
)

type CartImpl struct {
	cartRepo    domain.CartRepository
	cartService domain.CartService
}

func ProvideCart(cartRepo domain.CartRepository, cartService domain.CartService) *CartImpl {
	return &CartImpl{
		cartRepo:    cartRepo,
		cartService: cartService,
	}
}

var _ Cart = &CartImpl{}

func (c *CartImpl) Get(ctx context.Context, param GetCartParam) (*domain.Cart, error) {
	cart, err := c.cartRepo.Get(ctx, param.CartID)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (c *CartImpl) Create(ctx context.Context, param CreateCartParam) (*domain.Cart, error) {
	cart, err := c.cartService.Create(param.UserID)
	if err != nil {
		return nil, err
	}
	err = c.cartRepo.Save(ctx, cart)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (c *CartImpl) CreateItem(ctx context.Context, param CreateCartItemParam) (*domain.CartItem, error) {
	cart, err := c.cartRepo.Get(ctx, param.CartID)
	if err != nil {
		return nil, err
	}
	
	// Verify cart belongs to user
	if cart.UserID != param.UserID {
		return nil, domain.ErrForbidden
	}
	
	cartItem, err := c.cartService.CreateItem(param.Data.ProductVariantID, param.Data.Quantity)
	if err != nil {
		return nil, err
	}
	
	err = c.cartService.AddItem(cart, *cartItem)
	if err != nil {
		return nil, err
	}
	
	err = c.cartRepo.Save(ctx, cart)
	if err != nil {
		return nil, err
	}
	
	return cartItem, nil
}

func (c *CartImpl) UpdateItem(ctx context.Context, param UpdateCartItemParam) (*domain.CartItem, error) {
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
	
	err = c.cartRepo.Save(ctx, cart)
	if err != nil {
		return nil, err
	}
	
	// Find and return updated item
	if cart.Items != nil {
		for _, item := range *cart.Items {
			if item.ID == param.ItemID {
				return &item, nil
			}
		}
	}
	
	return nil, domain.ErrNotFound
}

func (c *CartImpl) DeleteItem(ctx context.Context, param DeleteCartItemParam) error {
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
	
	return c.cartRepo.Save(ctx, cart)
}
