package application

import (
	"context"

	"backend/internal/delivery/http"

	"github.com/google/uuid"
)

// CartCache defines the interface for cart caching operations
type CartCache interface {
	// GetCart retrieves a cached cart by ID
	GetCart(ctx context.Context, cartID uuid.UUID) (*http.CartResponseDto, error)

	// SetCart caches a cart with the specified TTL in seconds
	SetCart(ctx context.Context, cartID uuid.UUID, cart *http.CartResponseDto) error

	// InvalidateCart removes the cached cart by ID
	InvalidateCart(ctx context.Context, cartID uuid.UUID) error

	// InvalidateUserCart removes the cached cart for a specific user
	InvalidateUserCart(ctx context.Context, userID uuid.UUID) error
}
