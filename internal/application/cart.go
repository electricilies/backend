package application

import (
	"context"

	"backend/internal/delivery/http"
	"backend/internal/domain"

	"github.com/google/uuid"
)

type Cart struct {
	cartRepo    domain.CartRepository
	cartService domain.CartService
	cartCache   CartCache
	productRepo domain.ProductRepository
}

func ProvideCart(cartRepo domain.CartRepository, cartService domain.CartService, cartCache CartCache, productRepo domain.ProductRepository) *Cart {
	return &Cart{
		cartRepo:    cartRepo,
		cartService: cartService,
		cartCache:   cartCache,
		productRepo: productRepo,
	}
}

var _ http.CartApplication = (*Cart)(nil)

func (c *Cart) Get(ctx context.Context, param http.GetCartRequestDto) (*http.CartResponseDto, error) {
	cacheParam := CartCacheParam{ID: param.CartID}

	if cachedCart, err := c.cartCache.Get(ctx, cacheParam); err == nil {
		return cachedCart, nil
	}

	cart, err := c.cartRepo.Get(
		ctx,
		domain.CartRepositoryGetParam{
			ID: param.CartID,
		},
	)
	if err != nil {
		return nil, err
	}

	cartDto := http.ToCartResponseDto(cart)

	if err := c.enrichCartItems(ctx, cartDto, cart); err != nil {
		return nil, err
	}

	_ = c.cartCache.Set(ctx, cacheParam, cartDto)

	return cartDto, nil
}

func (c *Cart) GetByUser(ctx context.Context, param http.GetCartByUserRequestDto) (*http.CartResponseDto, error) {
	cart, err := c.cartRepo.Get(
		ctx,
		domain.CartRepositoryGetParam{
			UserID: param.UserID,
		},
	)
	if err != nil {
		return nil, err
	}

	cartDto := http.ToCartResponseDto(cart)

	if err := c.enrichCartItems(ctx, cartDto, cart); err != nil {
		return nil, err
	}

	return cartDto, nil
}

func (c *Cart) Create(ctx context.Context, param http.CreateCartRequestDto) (*http.CartResponseDto, error) {
	cart, err := domain.NewCart(param.Data.UserID)
	if err != nil {
		return nil, err
	}
	if err := c.cartService.Validate(*cart); err != nil {
		return nil, err
	}
	err = c.cartRepo.Save(ctx, domain.CartRepositorySaveParam{Cart: *cart})
	if err != nil {
		return nil, err
	}

	cartDto := http.ToCartResponseDto(cart)

	if err := c.enrichCartItems(ctx, cartDto, cart); err != nil {
		return nil, err
	}

	cacheParam := CartCacheParam{ID: cart.ID}
	_ = c.cartCache.Set(ctx, cacheParam, cartDto)

	return cartDto, nil
}

func (c *Cart) CreateItem(ctx context.Context, param http.CreateCartItemRequestDto) (*http.CartItemResponseDto, error) {
	cart, err := c.cartRepo.Get(
		ctx,
		domain.CartRepositoryGetParam{
			ID: param.CartID,
		},
	)
	if err != nil {
		return nil, err
	}

	// Verify cart belongs to user
	if cart.UserID != param.UserID {
		return nil, domain.ErrForbidden
	}

	cartItem, err := domain.NewCartItem(
		param.Data.ProductID,
		param.Data.ProductVariantID,
		param.Data.Quantity,
	)
	if err != nil {
		return nil, err
	}

	*cartItem = cart.UpsertItem(*cartItem)
	if err := c.cartService.Validate(*cart); err != nil {
		return nil, err
	}

	err = c.cartRepo.Save(ctx, domain.CartRepositorySaveParam{Cart: *cart})
	if err != nil {
		return nil, err
	}

	_ = c.cartCache.Invalidate(ctx, CartCacheParam{ID: param.CartID})

	// Enrich the cart item with product and variant data
	product, err := c.productRepo.Get(ctx, domain.ProductRepositoryGetParam{ProductID: cartItem.ProductID})
	if err != nil {
		return nil, err
	}

	variant := product.GetVariantByID(cartItem.ProductVariantID)
	if variant == nil {
		return nil, domain.ErrNotFound
	}

	return http.ToCartItemResponseDto(cartItem, product, variant), nil
}

func (c *Cart) UpdateItem(ctx context.Context, param http.UpdateCartItemRequestDto) (*http.CartItemResponseDto, error) {
	cart, err := c.cartRepo.Get(
		ctx,
		domain.CartRepositoryGetParam{
			ID: param.CartID,
		},
	)
	if err != nil {
		return nil, err
	}

	// Verify cart belongs to user
	if cart.UserID != param.UserID {
		return nil, domain.ErrForbidden
	}

	cart.UpdateItem(param.ItemID, param.Data.Quantity)
	if err := c.cartService.Validate(*cart); err != nil {
		return nil, err
	}

	err = c.cartRepo.Save(ctx, domain.CartRepositorySaveParam{Cart: *cart})
	if err != nil {
		return nil, err
	}

	// Find and return updated item
	if cart.Items != nil {
		for _, item := range cart.Items {
			if item.ID == param.ItemID {
				_ = c.cartCache.Invalidate(ctx, CartCacheParam{ID: param.CartID})

				// Enrich the cart item with product and variant data
				product, err := c.productRepo.Get(ctx, domain.ProductRepositoryGetParam{ProductID: item.ProductID})
				if err != nil {
					return nil, err
				}

				variant := product.GetVariantByID(item.ProductVariantID)
				if variant == nil {
					return nil, domain.ErrNotFound
				}

				return http.ToCartItemResponseDto(&item, product, variant), nil
			}
		}
	}

	return nil, domain.ErrNotFound
}

func (c *Cart) DeleteItem(ctx context.Context, param http.DeleteCartItemRequestDto) error {
	cart, err := c.cartRepo.Get(
		ctx,
		domain.CartRepositoryGetParam{
			ID: param.CartID,
		},
	)
	if err != nil {
		return err
	}

	// Verify cart belongs to user
	if cart.UserID != param.UserID {
		return domain.ErrForbidden
	}

	cart.RemoveItem(param.ItemID)
	if err := c.cartService.Validate(*cart); err != nil {
		return err
	}

	err = c.cartRepo.Save(ctx, domain.CartRepositorySaveParam{Cart: *cart})
	if err != nil {
		return err
	}

	_ = c.cartCache.Invalidate(ctx, CartCacheParam{ID: param.CartID})

	return nil
}

func (c *Cart) enrichCartItems(ctx context.Context, cartDto *http.CartResponseDto, cart *domain.Cart) error {
	if len(cart.Items) == 0 {
		return nil
	}

	// Collect unique product IDs
	productIDsMap := make(map[uuid.UUID]struct{})
	for _, item := range cart.Items {
		productIDsMap[item.ProductID] = struct{}{}
	}

	productIDs := make([]uuid.UUID, 0, len(productIDsMap))
	for id := range productIDsMap {
		productIDs = append(productIDs, id)
	}

	// Fetch all products at once
	products, err := c.productRepo.List(
		ctx,
		domain.ProductRepositoryListParam{
			IDs:     productIDs,
			Deleted: domain.DeletedExcludeParam,
		},
	)
	if err != nil {
		return err
	}

	// Create a map of product ID to product
	productMap := make(map[uuid.UUID]*domain.Product)
	for i := range *products {
		productMap[(*products)[i].ID] = &(*products)[i]
	}

	// Build enriched cart items
	enrichedItems := make([]http.CartItemResponseDto, 0, len(cart.Items))
	for _, item := range cart.Items {
		product, exists := productMap[item.ProductID]
		if !exists {
			continue
		}

		variant := product.GetVariantByID(item.ProductVariantID)
		if variant == nil {
			continue
		}

		itemDto := http.ToCartItemResponseDto(&item, product, variant)
		if itemDto != nil {
			enrichedItems = append(enrichedItems, *itemDto)
		}
	}

	cartDto.WithCartItems(enrichedItems)
	return nil
}
