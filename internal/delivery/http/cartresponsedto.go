package http

import (
	"time"

	"backend/internal/domain"

	"github.com/google/uuid"
)

// CartResponseDto represents the response structure for a cart
type CartResponseDto struct {
	ID        uuid.UUID             `json:"id"        binding:"required"`
	Items     []CartItemResponseDto `json:"items"     binding:"required"`
	UserID    uuid.UUID             `json:"userId"    binding:"required"`
	UpdatedAt time.Time             `json:"updatedAt" binding:"required"`
}

// CartItemResponseDto represents the response structure for a cart item
type CartItemResponseDto struct {
	ID             uuid.UUID                         `json:"id"             binding:"required"`
	Product        CartItemProductResponseDto        `json:"product"        binding:"required"`
	ProductVariant CartItemProductVariantResponseDto `json:"productVariant" binding:"required"`
	Quantity       int                               `json:"quantity"       binding:"required"`
}

type CartItemProductResponseDto struct {
	ID     uuid.UUID `json:"id"     binding:"required"`
	Name   string    `json:"name"   binding:"required"`
	Price  float64   `json:"price"  binding:"required"`
	Rating float64   `json:"rating" binding:"required"`
}

type CartItemProductVariantResponseDto struct {
	ID       uuid.UUID                 `json:"id"       binding:"required"`
	SKU      string                    `json:"sku"      binding:"required"`
	Price    int64                     `json:"price"    binding:"required"`
	Quantity int                       `json:"quantity" binding:"required"`
	Images   []ProductImageResponseDto `json:"images"   binding:"required"`
}

// ToCartResponseDto maps a domain.Cart to CartResponseDto
// Note: Items will be empty, use WithCartItems to populate with product data
func ToCartResponseDto(cart *domain.Cart) *CartResponseDto {
	if cart == nil {
		return nil
	}

	return &CartResponseDto{
		ID:        cart.ID,
		UserID:    cart.UserID,
		UpdatedAt: cart.UpdatedAt,
	}
}

// WithCartItems enriches CartResponseDto with cart items that have product and variant data
func (c *CartResponseDto) WithCartItems(items []CartItemResponseDto) *CartResponseDto {
	c.Items = items
	return c
}

// ToCartItemResponseDto creates a CartItemResponseDto from domain.CartItem with product and variant
func ToCartItemResponseDto(
	item *domain.CartItem,
	product *domain.Product,
	variant *domain.ProductVariant,
) *CartItemResponseDto {
	if item == nil {
		return nil
	}

	itemDto := &CartItemResponseDto{
		ID:       item.ID,
		Quantity: item.Quantity,
	}

	if product != nil {
		itemDto.Product = CartItemProductResponseDto{
			ID:     product.ID,
			Name:   product.Name,
			Price:  float64(product.Price),
			Rating: product.Rating,
		}
	}

	if variant != nil {
		images := make([]ProductImageResponseDto, 0, len(variant.Images))
		for _, img := range variant.Images {
			images = append(images, *ToProductImageResponseDto(&img))
		}

		itemDto.ProductVariant = CartItemProductVariantResponseDto{
			ID:       variant.ID,
			SKU:      variant.SKU,
			Price:    variant.Price,
			Quantity: variant.Quantity,
			Images:   images,
		}
	}

	return itemDto
}
